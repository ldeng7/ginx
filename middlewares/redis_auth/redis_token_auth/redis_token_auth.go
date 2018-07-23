package redis_token_auth

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ldeng7/ginx"
	"github.com/ldeng7/ginx/middlewares/redis_auth"
)

const RED_KEY_PREF = "tauth:"
const GIN_META_UID = "tauth:uid"

type RedisTokenAuth struct {
	*redis_auth.RedisAuth
}

func init() {
	rand.Seed(time.Now().Unix())
}

func New(red *redis.Client, namespace string) *RedisTokenAuth {
	return &RedisTokenAuth{
		redis_auth.New(red, namespace, RED_KEY_PREF),
	}
}

func (a *RedisTokenAuth) auth(gc *gin.Context) (int, string, error) {
	h := gc.Request.Header.Get("X-Access-Token")
	parts := strings.SplitN(h, ":", 2)
	if 2 != len(parts) {
		return http.StatusUnauthorized, "", errors.New("nauthorized")
	}
	uid, token := parts[0], parts[1]
	if 0 == len(uid) || 0 == len(token) {
		return http.StatusUnauthorized, "", errors.New("unauthorized")
	}
	status, err := a.Read(uid, token)
	return status, uid, err
}

func (a *RedisTokenAuth) Middleware() gin.HandlerFunc {
	return func(gc *gin.Context) {
		status, uid, err := a.auth(gc)
		if http.StatusOK != status {
			c := ginx.Context{gc}
			c.RenderError(&ginx.RespError{StatusCode: status, Message: err.Error()})
			gc.Abort()
			return
		}
		gc.Set(GIN_META_UID, uid)
		gc.Next()
	}
}

func (a *RedisTokenAuth) Set(uid string, ttl time.Duration) (string, error) {
	h := md5.New()
	h.Write([]byte(time.Now().String()))
	h.Write([]byte(strconv.Itoa(rand.Int())))
	token := fmt.Sprintf("%x", h.Sum(nil))

	if err := a.Write(uid, token, ttl); nil != err {
		return "", err
	}
	return token, nil
}

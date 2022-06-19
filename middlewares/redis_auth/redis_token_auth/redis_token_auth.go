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
	"github.com/go-redis/redis/v8"
	"github.com/ldeng7/ginx/ginx"
	"github.com/ldeng7/ginx/middlewares/redis_auth"
)

const RED_KEY_PREF = "tauth:"
const GIN_META_UID = "tauth:uid"

type RedisTokenAuth struct {
	*redis_auth.RedisAuth
	r *rand.Rand
}

func New(red *redis.Client, namespace string) *RedisTokenAuth {
	return &RedisTokenAuth{
		RedisAuth: redis_auth.New(red, namespace, RED_KEY_PREF),
		r:         rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (a *RedisTokenAuth) auth(gc *gin.Context) (int, string, error) {
	h := gc.Request.Header.Get("X-Access-Token")
	parts := strings.SplitN(h, ":", 2)
	if len(parts) != 2 {
		return http.StatusUnauthorized, "", errors.New("nauthorized")
	}
	uid, token := parts[0], parts[1]
	if len(uid) == 0 || len(token) == 0 {
		return http.StatusUnauthorized, "", errors.New("unauthorized")
	}
	status, err := a.Read(uid, token)
	return status, uid, err
}

func (a *RedisTokenAuth) Middleware() gin.HandlerFunc {
	return func(gc *gin.Context) {
		status, uid, err := a.auth(gc)
		if http.StatusOK != status {
			g := ginx.G{Context: gc}
			g.RenderError(&ginx.RespError{Status: status, Message: err.Error()})
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
	h.Write([]byte(strconv.Itoa(a.r.Int())))
	token := fmt.Sprintf("%x", h.Sum(nil))

	if err := a.Write(uid, token, ttl); nil != err {
		return "", err
	}
	return token, nil
}

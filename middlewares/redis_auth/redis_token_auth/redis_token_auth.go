package redis_token_auth

import (
	"crypto/md5"
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
	"github.com/ldeng7/go-x/logx"
)

const RED_KEY_PREF = "tauth:"
const GIN_META_UID = "tauth:uid"

type RedisTokenAuth struct {
	*redis_auth.RedisAuth
}

func init() {
	rand.Seed(time.Now().Unix())
}

func New(red *redis.Client, namespace string, logger *logx.Logger) *RedisTokenAuth {
	return &RedisTokenAuth{
		redis_auth.New(red, namespace, RED_KEY_PREF, logger),
	}
}

func (a *RedisTokenAuth) auth(c *gin.Context) (int, string) {
	h := c.Request.Header.Get("X-Access-Token")
	parts := strings.SplitN(h, ":", 2)
	if 2 != len(parts) {
		return http.StatusUnauthorized, ""
	}
	uid, token := parts[0], parts[1]
	if 0 == len(uid) || 0 == len(token) {
		return http.StatusUnauthorized, ""
	}
	return a.Read(uid, token), uid
}

func (a *RedisTokenAuth) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, uid := a.auth(c)
		if http.StatusOK != status {
			ginx.RenderError(c, &ginx.RespError{StatusCode: status})
			c.Abort()
			return
		}
		c.Set(GIN_META_UID, uid)
		c.Next()
	}
}

func (a *RedisTokenAuth) Set(uid string, domain string, ttl time.Duration) (string, error) {
	h := md5.New()
	h.Write([]byte(time.Now().String()))
	h.Write([]byte(strconv.Itoa(rand.Int())))
	token := fmt.Sprintf("%x", h.Sum(nil))

	key := fmt.Sprintf("%s%s@%s", a.Prefix, uid, domain)
	if err := a.Red.Set(key, token, ttl).Err(); nil != err {
		a.Logger.Err(err.Error())
		return "", err
	}
	return token, nil
}

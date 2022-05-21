package redis_basic_auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ldeng7/ginx"
	"github.com/ldeng7/ginx/middlewares/redis_auth"
)

const RED_KEY_PREF = "bauth:"
const GIN_META_UID = "bauth:uid"

type RedisBasicAuth struct {
	*redis_auth.RedisAuth
	realm string
}

func New(red *redis.Client, namespace, realm string) *RedisBasicAuth {
	return &RedisBasicAuth{
		RedisAuth: redis_auth.New(red, namespace, RED_KEY_PREF),
		realm:     realm,
	}
}

func decodeAuth(s string) (string, string) {
	if !strings.HasPrefix(s, "Basic ") {
		return "", ""
	}
	s = s[6:]
	if len(s) == 0 {
		return "", ""
	}

	bs, err := base64.StdEncoding.DecodeString(s)
	if nil != err {
		return "", ""
	}
	s = string(bs)

	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func (a *RedisBasicAuth) auth(gc *gin.Context) (int, string, error) {
	u, p := decodeAuth(gc.GetHeader("Authorization"))
	if len(u) == 0 || len(p) == 0 {
		return http.StatusUnauthorized, "", errors.New("unauthorized")
	}
	status, err := a.Read(u, p)
	return status, u, err
}

func (a *RedisBasicAuth) Middleware() gin.HandlerFunc {
	return func(gc *gin.Context) {
		status, uid, err := a.auth(gc)
		c := ginx.Context{Context: gc}
		if http.StatusOK != status {
			if http.StatusUnauthorized == status {
				c.Header("WWW-Authenticate", fmt.Sprintf(`Basic Realm="%s"`, a.realm))
			}
			c.RenderError(&ginx.RespError{Status: status, Message: err.Error()})
			c.Abort()
			return
		}
		c.Set(GIN_META_UID, uid)
		c.Next()
	}
}

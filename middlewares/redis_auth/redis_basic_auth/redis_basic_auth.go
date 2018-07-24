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
	if 0 == len(s) {
		return "", ""
	}

	bs, err := base64.StdEncoding.DecodeString(s)
	if nil != err {
		return "", ""
	}
	s = string(bs)

	parts := strings.Split(s, ":")
	if 2 != len(parts) {
		return "", ""
	}
	return parts[0], parts[1]
}

func (a *RedisBasicAuth) auth(gc *gin.Context) (int, string, error) {
	u, p := decodeAuth(gc.GetHeader("Authorization"))
	if 0 == len(u) || 0 == len(p) {
		return http.StatusUnauthorized, "", errors.New("unauthorized")
	}
	status, err := a.Read(u, p)
	return status, u, err
}

func (a *RedisBasicAuth) Middleware() gin.HandlerFunc {
	return func(gc *gin.Context) {
		status, uid, err := a.auth(gc)
		if http.StatusOK != status {
			if http.StatusUnauthorized == status {
				c.Header("WWW-Authenticate", fmt.Sprintf(`Basic Realm="%s"`, a.realm))
			}
			c := ginx.Context{gc}
			c.RenderError(&ginx.RespError{Status: status, Message: err.Error()})
			gc.Abort()
			return
		}
		gc.Set(GIN_META_UID, uid)
		gc.Next()
	}
}

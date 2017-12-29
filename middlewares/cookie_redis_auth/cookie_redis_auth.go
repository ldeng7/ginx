package cookie_redis_auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const COOKIE_NAME_UID = "uid"
const COOKIE_NAME_TOKEN = "token"
const GIN_META_UID = "uid"
const RED_KEY_PREF = "cooauth:"

type CookieRedisAuth struct {
	red    *redis.Client
	prefix string
	ttl    time.Duration
}

func New(red *redis.Client, namespace string, ttl time.Duration) *CookieRedisAuth {
	prefix := RED_KEY_PREF
	if len(namespace) > 0 {
		prefix = namespace + ":" + prefix
	}
	return &CookieRedisAuth{
		red:    red,
		prefix: prefix,
		ttl:    ttl,
	}
}

func (c *CookieRedisAuth) auth(ctx *gin.Context) (int, *http.Cookie) {
	cookieUid, err := ctx.Request.Cookie(COOKIE_NAME_UID)
	if nil != err {
		return http.StatusUnauthorized, nil
	}
	cookieToken, err := ctx.Request.Cookie(COOKIE_NAME_TOKEN)
	if nil != err {
		return http.StatusUnauthorized, nil
	}
	token, err := c.red.Get(c.prefix + cookieUid.Value).Result()
	if redis.Nil == err {
		return http.StatusUnauthorized, nil
	} else if nil != err {
		return http.StatusInternalServerError, nil
	} else if token != cookieToken.Value {
		return http.StatusUnauthorized, nil
	}
	return http.StatusOK, cookieUid
}

func (c *CookieRedisAuth) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status, cookieUid := c.auth(ctx)
		if http.StatusOK != status {
			ctx.AbortWithStatus(status)
			return
		}
		ctx.Set(GIN_META_UID, cookieUid.Value)
		ctx.Next()
	}
}

func (c *CookieRedisAuth) Set(ctx *gin.Context, uid, token string) error {
	now := time.Now()
	if err := c.red.Set(c.prefix+uid, token, c.ttl).Err(); nil != err {
		return err
	}
	expire := now.Add(c.ttl)
	maxAge := int(c.ttl / time.Second)
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    COOKIE_NAME_UID,
		Value:   uid,
		Path:    "/",
		Expires: expire,
		MaxAge:  maxAge,
	})
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    COOKIE_NAME_TOKEN,
		Value:   token,
		Path:    "/",
		Expires: expire,
		MaxAge:  maxAge,
	})
	return nil
}

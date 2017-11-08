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
const REDIS_KEY = "token:"

type CookieRedisAuth struct {
	red    *redis.Client
	prefix string
	ttl    time.Duration
	path   string
}

func New(red *redis.Client, namespace string, ttl time.Duration, path string) *CookieRedisAuth {
	prefix := "cooauth:"
	if len(namespace) > 0 {
		prefix = namespace + ":" + prefix
	}
	return &CookieRedisAuth{
		red:    red,
		prefix: prefix,
		ttl:    ttl,
		path:   path,
	}
}

func (c *CookieRedisAuth) redisKey(uid string) string {
	return c.prefix + uid
}

func (c *CookieRedisAuth) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookieUid, err := ctx.Request.Cookie(COOKIE_NAME_UID)
		if nil != err {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		cookieToken, err := ctx.Request.Cookie(COOKIE_NAME_TOKEN)
		if nil != err {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		token, err := c.red.Get(c.redisKey(cookieUid.Value)).Result()
		if redis.Nil == err {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		} else if nil != err {
			ctx.Status(http.StatusInternalServerError)
			ctx.Abort()
			return
		} else if token != cookieToken.Value {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set(GIN_META_UID, cookieUid.Value)
		ctx.Next()
	}
}

func (c *CookieRedisAuth) Set(ctx *gin.Context, uid, token string) error {
	now := time.Now()
	if err := c.red.Set(c.redisKey(uid), token, c.ttl).Err(); nil != err {
		return err
	}
	expire := now.Add(c.ttl)
	maxAge := int(c.ttl / time.Second)
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    COOKIE_NAME_UID,
		Value:   uid,
		Path:    c.path,
		Expires: expire,
		MaxAge:  maxAge,
	})
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    COOKIE_NAME_TOKEN,
		Value:   token,
		Path:    c.path,
		Expires: expire,
		MaxAge:  maxAge,
	})
	return nil
}

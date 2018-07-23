package redis_auth

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/ldeng7/go-logx/logx"
)

type RedisAuth struct {
	Red    *redis.Client
	Prefix string
	Logger *logx.Logger
}

func New(red *redis.Client, namespace, prefix string, logger *logx.Logger) *RedisAuth {
	if len(namespace) > 0 {
		prefix = namespace + ":" + prefix
	}
	return &RedisAuth{
		Red:    red,
		Prefix: prefix,
		Logger: logger,
	}
}

func (a *RedisAuth) Read(uid, token string) int {
	tokenRed, err := a.Red.Get(a.Prefix + uid).Result()
	if redis.Nil == err {
		return http.StatusUnauthorized
	} else if nil != err {
		a.Logger.Err(err.Error())
		return http.StatusInternalServerError
	} else if tokenRed != token {
		return http.StatusUnauthorized
	}
	return http.StatusOK
}

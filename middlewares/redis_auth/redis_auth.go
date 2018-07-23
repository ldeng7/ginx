package redis_auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

type RedisAuth struct {
	Red    *redis.Client
	Prefix string
}

func New(red *redis.Client, namespace, prefix string) *RedisAuth {
	if len(namespace) > 0 {
		prefix = namespace + ":" + prefix
	}
	return &RedisAuth{
		Red:    red,
		Prefix: prefix,
	}
}

func (a *RedisAuth) Read(uid, token string) (int, error) {
	tokenRed, err := a.Red.Get(a.Prefix + uid).Result()
	if redis.Nil == err {
		return http.StatusUnauthorized, errors.New("unauthorized")
	} else if nil != err {
		return http.StatusInternalServerError, err
	} else if tokenRed != token {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}
	return http.StatusOK, nil
}

func (a *RedisAuth) Write(uid, token string, ttl time.Duration) error {
	return a.Red.Set(a.Prefix+uid, token, ttl).Err()
}

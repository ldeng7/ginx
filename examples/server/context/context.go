package context

import (
	"github.com/go-redis/redis/v8"
	"github.com/ldeng7/ginx/ginx"
	"github.com/ldeng7/go-logger-lite/logger"
	"gorm.io/gorm"
)

type Config struct {
	GinMode  string `yaml:"gin_mode"`
	LogPath  string `yaml:"log_path"`
	LogLevel int    `yaml:"log_level"`
	HttpPort int    `yaml:"http_port"`

	Mysql *ginx.GormMysqlConf
	Redis *ginx.RedisConf
}

type Context struct {
	Config *Config
	Logger *logger.Logger
	Gorm   *gorm.DB
	Red    *redis.Client
}

var singleton *Context

func Instance() *Context {
	return singleton
}

func Init(confPath string) error {
	var err error
	ctx := &Context{}
	singleton = ctx

	conf := &Config{}
	if err = ginx.InitConf(confPath, conf); nil != err {
		return err
	}
	ctx.Config = conf

	if ctx.Logger, err = ginx.NewLogger(nil, conf.LogPath, conf.LogLevel); nil != err {
		return err
	}

	if ctx.Gorm, err = ginx.NewGormMysql(conf.Mysql); nil != err {
		return err
	}

	ctx.Red = ginx.NewRedis(conf.Redis)

	return nil
}

package ginx

import (
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ldeng7/go-logger-lite/logger"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConf(path string, conf any) error {
	confStr, err := ioutil.ReadFile(path)
	if nil != err {
		return err
	}
	if err := yaml.Unmarshal(confStr, conf); nil != err {
		return err
	}
	return nil
}

func NewLogger(writer io.Writer, logPath string, logLevel int) (*logger.Logger, error) {
	l, err := logger.New(&logger.InitArgs{
		Writer:   writer,
		Filename: logPath + ".error.log",
		Flags:    log.LstdFlags | log.Lshortfile,
		LogLevel: logLevel,
	})
	if err != nil {
		return nil, err
	}

	la, err := logger.New(&logger.InitArgs{
		Writer:   writer,
		Filename: logPath + ".access.log",
	})
	if err != nil {
		return nil, err
	}
	gin.DefaultWriter = la.GetWriter()

	return l, nil
}

type GormMysqlConf struct {
	Dsn     string
	Mysql   *mysql.Config
	Gorm    *gorm.Config
	MaxIdle int
	MaxOpen int
}

func NewGormMysql(conf *GormMysqlConf) (*gorm.DB, error) {
	var dial gorm.Dialector
	if nil == conf.Mysql {
		dial = mysql.Open(conf.Dsn)
	} else {
		dial = mysql.New(*conf.Mysql)
	}
	db, err := gorm.Open(dial, conf.Gorm)
	if nil != err {
		return nil, err
	}

	sqlDb, _ := db.DB()
	if conf.MaxIdle > 0 {
		sqlDb.SetMaxIdleConns(conf.MaxIdle)
	}
	if conf.MaxOpen > 0 {
		sqlDb.SetMaxOpenConns(conf.MaxOpen)
	}

	return db, nil
}

type RedisInstConf struct {
	Addr           string
	Password       string
	DialTimeoutMs  int64 `yaml:"dial_timeout_ms"`
	ReadTimeoutMs  int64 `yaml:"read_timeout_ms"`
	WriteTimeoutMs int64 `yaml:"write_timeout_ms"`
}

type RedisConf struct {
	Red *RedisInstConf
	Db  int
}

func NewRedis(conf *RedisConf) *redis.Client {
	red := redis.NewClient(&redis.Options{
		Addr:         conf.Red.Addr,
		Password:     conf.Red.Password,
		DB:           conf.Db,
		DialTimeout:  time.Duration(conf.Red.DialTimeoutMs) * time.Millisecond,
		ReadTimeout:  time.Duration(conf.Red.ReadTimeoutMs) * time.Millisecond,
		WriteTimeout: time.Duration(conf.Red.WriteTimeoutMs) * time.Millisecond,
	})
	return red
}

package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitGorm(dsn string, maxIdle, maxOpen int) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if nil != err {
		return nil, err
	}
	if maxIdle <= 0 {
		maxIdle = 8
	}
	if maxOpen <= 0 {
		maxOpen = 128
	}
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.LogMode(gin.ReleaseMode != gin.Mode())
	return db, nil
}

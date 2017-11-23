package recovery_lite

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/go-x/logx"
)

func recovery(c *gin.Context, logger *logx.Logger, callback func(p interface{})) {
	p := recover()
	if nil == p {
		return
	}

	if nil != callback {
		callback(p)
	}
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file, line = "?", 0
	}
	logger.Errf("panic on %s:%d:", file, line)
	logger.Err(p)
	c.AbortWithStatus(500)
}

func Recovery(logger *logx.Logger, callback func(p interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer recovery(c, logger, callback)
		c.Next()
	}
}

package recovery_lite

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/ginx"
	"github.com/ldeng7/go-x/logx"
)

func recovery(c *gin.Context, logger *logx.Logger, callback func(*gin.Context, interface{})) {
	p := recover()
	if nil == p {
		return
	}

	if nil != callback {
		callback(c, p)
	}
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file, line = "?", 0
	}
	logger.Errf("panic on %s:%d:", file, line)
	logger.Err(p)
	ginx.RenderError(c, &ginx.RespError{StatusCode: http.StatusInternalServerError})
	c.Abort()
}

func Recovery(logger *logx.Logger, callback func(*gin.Context, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer recovery(c, logger, callback)
		c.Next()
	}
}

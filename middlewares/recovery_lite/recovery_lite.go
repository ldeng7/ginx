package recovery_lite

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/ginx"
	"github.com/ldeng7/go-logx/logx"
)

func recovery(gc *gin.Context, logger *logx.Logger, callback func(*gin.Context, interface{})) {
	p := recover()
	if nil == p {
		return
	}

	if nil != callback {
		callback(gc, p)
	}
	logger.Err("panic: ", p)
	for i := 3; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		logger.Errf("panic on %s:%d", file, line)
	}

	c := ginx.Context{gc}
	c.RenderError(&ginx.RespError{Status: http.StatusInternalServerError})
	gc.Abort()
}

func Recovery(logger *logx.Logger, callback func(*gin.Context, interface{})) gin.HandlerFunc {
	return func(gc *gin.Context) {
		defer recovery(gc, logger, callback)
		gc.Next()
	}
}

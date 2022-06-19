package recovery_lite

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/ginx/ginx"
	"github.com/ldeng7/go-logger-lite/logger"
)

func recovery(gc *gin.Context, logger *logger.Logger, callback func(*gin.Context, any)) {
	p := recover()
	if nil == p {
		return
	}

	if nil != callback {
		callback(gc, p)
	}
	logger.Err("panic: ", p)
	for i := 2; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		logger.Errf("panic on %s:%d", file, line)
	}

	g := ginx.G{Context: gc}
	g.RenderError(&ginx.RespError{Status: http.StatusInternalServerError})
	gc.Abort()
}

func Recovery(logger *logger.Logger, callback func(*gin.Context, any)) gin.HandlerFunc {
	return func(gc *gin.Context) {
		defer recovery(gc, logger, callback)
		gc.Next()
	}
}

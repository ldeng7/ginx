package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var metaRouteRegisters map[string]func(r *gin.Engine) = map[string]func(r *gin.Engine){
	"noroute": func(r *gin.Engine) {
		r.NoRoute(func(c *gin.Context) {
			RenderError(c, &RespError{StatusCode: http.StatusNotFound})
		})
	},

	"health": func(r *gin.Engine) {
		r.GET("/monitors/health", func(c *gin.Context) {
			RenderData(c, nil)
		})
	},
}

func MetaRouteRegister(r *gin.Engine, strs ...string) {
	for _, str := range strs {
		if register, ok := metaRouteRegisters[str]; ok {
			register(r)
		}
	}
}

func MetaRouteRegisterAll(r *gin.Engine) {
	for _, register := range metaRouteRegisters {
		register(r)
	}
}

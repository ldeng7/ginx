package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var metaRoutes map[string]func(r *gin.Engine) = map[string]func(r *gin.Engine){
	"noroute": func(r *gin.Engine) {
		r.NoRoute(func(gc *gin.Context) {
			c := &Context{gc}
			c.RenderError(&RespError{Status: http.StatusNotFound})
		})
	},

	"health": func(r *gin.Engine) {
		r.GET("/monitors/health", func(gc *gin.Context) {
			c := &Context{gc}
			c.RenderData(nil)
		})
	},
}

func MetaRouteRegister(r *gin.Engine, strs ...string) {
	for _, str := range strs {
		if register, ok := metaRoutes[str]; ok {
			register(r)
		}
	}
}

func MetaRouteRegisterAll(r *gin.Engine) {
	for _, register := range metaRoutes {
		register(r)
	}
}

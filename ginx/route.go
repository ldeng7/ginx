package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICreateController interface {
	Create(*G)
}

type IListController interface {
	List(*G)
}

type IGetController interface {
	Get(*G)
}

type IUpdateController interface {
	Update(*G)
}

type IDeleteController interface {
	Delete(*G)
}

func AddRoutes(group *gin.RouterGroup, path, pluralPath string, controller any) {
	if con, ok := controller.(ICreateController); ok {
		group.POST(path, func(gc *gin.Context) {
			con.Create(&G{Context: gc})
		})
	}
	if con, ok := controller.(IListController); ok {
		group.GET(pluralPath, func(gc *gin.Context) {
			con.List(&G{Context: gc})
		})
	}
	if con, ok := controller.(IGetController); ok {
		group.GET(path+"/:id", func(gc *gin.Context) {
			con.Get(&G{Context: gc})
		})
	}
	if con, ok := controller.(IUpdateController); ok {
		group.PUT(path+"/:id", func(gc *gin.Context) {
			con.Update(&G{Context: gc})
		})
	}
	if con, ok := controller.(IDeleteController); ok {
		group.DELETE(path+"/:id", func(gc *gin.Context) {
			con.Delete(&G{Context: gc})
		})
	}
}

var metaRoutes map[string]func(r *gin.Engine) = map[string]func(r *gin.Engine){
	"noroute": func(r *gin.Engine) {
		r.NoRoute(func(gc *gin.Context) {
			g := &G{Context: gc}
			g.RenderError(&RespError{Status: http.StatusNotFound})
		})
	},

	"health": func(r *gin.Engine) {
		r.GET("/monitors/health", func(gc *gin.Context) {
			g := &G{Context: gc}
			g.RenderData(nil)
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

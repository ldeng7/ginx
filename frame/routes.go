package frame

import (
	"github.com/gin-gonic/gin"
)

type IIndexController interface {
	Index(*gin.Context)
}

type ICreateController interface {
	Create(*gin.Context)
}

type IShowController interface {
	Show(*gin.Context)
}

type IUpdateController interface {
	Update(*gin.Context)
}

type IDestroyController interface {
	Destroy(*gin.Context)
}

func AddRoutes(group *gin.RouterGroup, path string, controller interface{}) {
	if c, ok := controller.(IIndexController); ok {
		group.GET(path, func(ctx *gin.Context) {
			c.Index(ctx)
		})
	}
	if c, ok := controller.(ICreateController); ok {
		group.POST(path, func(ctx *gin.Context) {
			c.Create(ctx)
		})
	}
	if c, ok := controller.(IShowController); ok {
		group.GET(path+"/:id", func(ctx *gin.Context) {
			c.Show(ctx)
		})
	}
	if c, ok := controller.(IUpdateController); ok {
		group.PUT(path+"/:id", func(ctx *gin.Context) {
			c.Update(ctx)
		})
	}
	if c, ok := controller.(IDestroyController); ok {
		group.DELETE(path+"/:id", func(ctx *gin.Context) {
			c.Destroy(ctx)
		})
	}
}

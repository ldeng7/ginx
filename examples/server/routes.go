package main

import (
	"github.com/ldeng7/ginx/examples/server/context"
	"github.com/ldeng7/ginx/examples/server/controllers"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/ginx/ginx"
	"github.com/ldeng7/ginx/middlewares/recovery_lite"
)

func setRoutes() *gin.Engine {
	ctx := context.Instance()
	gin.SetMode(ctx.Config.GinMode)
	r := gin.New()
	r.Use(gin.Logger(), recovery_lite.Recovery(ctx.Logger, nil))
	ginx.MetaRouteRegisterAll(r)

	v1Group := r.Group("v1")
	ginx.AddRoutes(v1Group, "user", "users", &controllers.UserController{})

	return r
}

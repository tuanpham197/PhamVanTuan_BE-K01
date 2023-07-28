package web_app

import (
	"fmt"

	"net/http/pprof"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/service"
	v1 "github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/v1"
)

type WebController struct {
	service.WebService
	Port int
}

func (c WebController) Run() {
	r := gin.Default()
	v1Router := r.Group("v1")
	v1.AddUserRouter(v1Router, &c.WebService)
	v1.AddPostRouter(v1Router, &c.WebService)

	initSwagger(r)

	r.GET("/debug/pprof/", func(context *gin.Context) {
		pprof.Index(context.Writer, context.Request)
	})

	r.GET("/debug/pprof/profile", func(context *gin.Context) {
		pprof.Profile(context.Writer, context.Request)
	})

	r.GET("/debug/pprof/trace", func(context *gin.Context) {
		pprof.Trace(context.Writer, context.Request)
	})

	r.Run(fmt.Sprintf(":%d", c.Port))
}

func initSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

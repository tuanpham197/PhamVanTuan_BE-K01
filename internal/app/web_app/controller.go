package web_app

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

	r.Run(fmt.Sprintf(":%d", c.Port))
}

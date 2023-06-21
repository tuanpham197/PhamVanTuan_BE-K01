package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/service"
)

func AddUserRouter(r *gin.RouterGroup, svc *service.WebService) {
	userRouter := r.Group("users")
	userRouter.GET("", svc.CheckUserNamePassword)
	userRouter.POST("", svc.CreateUser)
	userRouter.PUT("", svc.EditUser)
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/service"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/v1/middleware"
)

func AddUserRouter(r *gin.RouterGroup, svc *service.WebService) {
	userRouter := r.Group("users")
	// userRouter.Use(middleware.AuthenticateMiddleware)
	userRouter.POST("/check", svc.CheckUserNamePassword)
	userRouter.POST("", svc.CreateUser)
	userRouter.PUT("", svc.EditUser)

	// follow
	FriendRouter := r.Group("friends")
	FriendRouter.Use(middleware.AuthenticateMiddleware)
	FriendRouter.GET("/:user_id", svc.GetListFollower)
	FriendRouter.POST("/:user_id", svc.FollowUser)
	FriendRouter.DELETE("/:user_id", svc.UnFollowUser)
	FriendRouter.GET("/:user_id/posts", svc.GetPostList)
}

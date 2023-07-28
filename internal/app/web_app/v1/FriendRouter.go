package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/service"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/v1/middleware"
)

func FriendRouter(r *gin.RouterGroup, svc *service.WebService) {
	FriendRouter := r.Group("friends")

	FriendRouter.Use(middleware.AuthenticateMiddleware)
	FriendRouter.GET("/:user_id", svc.GetListFollower)
	FriendRouter.POST("/:user_id", svc.FollowUser)
	FriendRouter.DELETE("/:user_id", svc.UnFollowUser)
	FriendRouter.PUT("/:user_id/posts", svc.GetPostList)

}

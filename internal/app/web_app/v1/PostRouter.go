package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/service"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/app/web_app/v1/middleware"
)

func AddPostRouter(r *gin.RouterGroup, svc *service.WebService) {
	PostRouter := r.Group("posts")

	PostRouter.Use(middleware.AuthenticateMiddleware)
	PostRouter.GET("/", svc.GetPostList)
	PostRouter.POST("/", svc.CreatePost)
	PostRouter.GET("/:post_id", svc.GetPostDetail)
	PostRouter.PUT("/:post_id", svc.UpdatePost)
	PostRouter.DELETE("/:post_id", svc.DeletePost)
	PostRouter.POST("/:post_id/comments", svc.CommentPost)
	PostRouter.POST("/:user_id/likes", svc.LikePost)

}

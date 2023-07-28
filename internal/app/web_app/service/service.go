package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/configs"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/pkg/types"
	authen_and_post2 "github.com/tuanpham197/PhamVanTuan_BE-K01/pkg/client/authen_and_post"
	newsfeed2 "github.com/tuanpham197/PhamVanTuan_BE-K01/pkg/client/newsfeed"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/pkg/types/proto/pb/newsfeed"
)

type WebService struct {
	authenticateAndPostClient authen_and_post.AuthenticateAndPostClient
	newsfeedClient            newsfeed.NewsfeedClient
}

func NewWebService(conf *configs.WebConfig) (*WebService, error) {
	aapClient, err := authen_and_post2.NewClient(conf.AuthenticateAndPost.Hosts)
	if err != nil {
		return nil, err
	}
	newsfeedClient, err := newsfeed2.NewClient(conf.Newsfeed.Hosts)
	if err != nil {
		return nil, err
	}
	return &WebService{
		authenticateAndPostClient: aapClient,
		newsfeedClient:            newsfeedClient,
	}, nil
}

// GetBooks             godoc
// @Summary      Get books array
// @Description  Responds with the list of all books as JSON.
// @Param 		 request body types.LoginRequest true "login param"
// @Tags         books
// @Produce      json
// @Success		 200	{object} types.MessageResponse
// @Failure		 400	{object} types.MessageResponse
// @Failure		 500	{object} types.MessageResponse
// @Router		 /users/check [post]
func (svc *WebService) CheckUserNamePassword(ctx *gin.Context) {
	var jsonRequest types.LoginRequest
	err := ctx.ShouldBindJSON(&jsonRequest)

	if err != nil {
		ctx.JSON(http.StatusOK, &types.MessageResponse{Message: err.Error()})
		return
	}
	authentication, err := svc.authenticateAndPostClient.CheckUserAuthentication(ctx, &authen_and_post.UserInfo{
		UserName:     jsonRequest.UserName,
		UserPassword: jsonRequest.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, &types.MessageResponse{Message: err.Error()})
		return
	}
	if authentication.GetStatus() == authen_and_post.UserStatus_OK {
		// change this later
		ctx.SetCookie("session_id", fmt.Sprintf("%d", authentication.Info.UserId), 0, "", "", false, false)

		ctx.JSON(http.StatusOK, &types.MessageResponse{Message: "ok"})

	}

}

func (svc *WebService) CreateUser(ctx *gin.Context) {

	username := ctx.PostForm("username")
	// password := ctx.PostForm("password")
	lastname := ctx.PostForm("lastname")
	firstname := ctx.PostForm("firstname")
	dob := ctx.GetInt64("dob")
	email := ctx.PostForm("email")

	result, err := svc.authenticateAndPostClient.CreateUser(ctx, &authen_and_post.UserDetailInfo{
		UserName:  username,
		FirstName: firstname,
		LastName:  lastname,
		Dob:       int64(dob),
		Email:     email,
	})

	if err != nil {
		panic("Err Create user")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "crete successed",
		"user":    result.Info,
	})

}

func (svc *WebService) EditUser(ctx *gin.Context) {
	var userInfoRequest types.UpdateUserRequest
	err := ctx.ShouldBindJSON(&userInfoRequest)

	if err != nil {
		panic("Err parse json")
	}

	cookie, _ := ctx.Cookie("session_id")
	userId, _ := strconv.ParseInt(cookie, 10, 64)

	userInfoDetail, errEdit := svc.authenticateAndPostClient.EditUser(ctx, &authen_and_post.UserDetailInfo{
		LastName:  userInfoRequest.LastName,
		FirstName: userInfoRequest.FirstName,
		BirthDay:  userInfoRequest.Birthday,
		Password:  userInfoRequest.Password,
		UserId:    userId,
	})

	if errEdit != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "EDIT FAIL",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"user":    userInfoDetail,
	})
}

// See follow list: GET v1/friends/user_id  {} {users}
// Follow: POST v1/friends/user_id {msg}
// Unfollow: DELETE v1/friends/user_id {msg}
// See user posts GET v1/friends/user_id/posts {posts}

func (svc *WebService) GetListFollower(ctx *gin.Context) {
	idParam := ctx.Param("user_id")
	userID, _ := strconv.ParseInt(idParam, 10, 64)

	followers, err := svc.authenticateAndPostClient.GetUserFollower(ctx, &authen_and_post.UserInfo{
		UserId: userID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "GET FAIL",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"data":    followers,
	})
}

func (svc *WebService) FollowUser(ctx *gin.Context) {
	userIdStr := ctx.Param("user_id")
	cookie, _ := ctx.Cookie("session_id")
	currentUserId, _ := strconv.ParseInt(cookie, 10, 64)
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	result, err := svc.authenticateAndPostClient.FollowUser(ctx, &authen_and_post.UserFollowRequest{
		UserId:   currentUserId,
		FriendId: userId,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "FOLLOW FAIL",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"code":    http.StatusOK,
		"success": result,
	})
}

func (svc *WebService) UnFollowUser(ctx *gin.Context) {
	userIdStr := ctx.Param("user_id")
	cookie, _ := ctx.Cookie("session_id")
	currentUserId, _ := strconv.ParseInt(cookie, 10, 64)
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	_, err := svc.authenticateAndPostClient.UnFollowUser(ctx, &authen_and_post.UserFollowRequest{
		UserId:   currentUserId,
		FriendId: userId,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "UNFOLLOW FAIL",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "unfollow success",
	})
}

func (svc *WebService) GetPostList(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	userIDInt, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get post FAIL",
			"error":   err,
		})
		return
	}

	result, err := svc.authenticateAndPostClient.GetPostFriend(ctx, &authen_and_post.UserInfo{
		UserId: userIDInt,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get post FAIL",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})

}

func (svc *WebService) CreatePost(ctx *gin.Context) {
	var postRequest types.PostRequest
	err := ctx.ShouldBindJSON(&postRequest)

	if err != nil {
		ctx.JSON(http.StatusOK, &types.MessageResponse{Message: err.Error()})
		return
	}

	cookie, _ := ctx.Cookie("session_id")
	currentUserId, _ := strconv.ParseInt(cookie, 10, 64)

	result, err := svc.authenticateAndPostClient.CreatePost(ctx, &authen_and_post.PostRequest{
		UserId:           currentUserId,
		ContentText:      postRequest.ContextText,
		ContentImagePath: postRequest.ContextImagePath,
		Visible:          true,
		CreatedAt:        time.Now().GoString(),
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Create post fail",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create success",
		"data":    result,
	})
}

func (svc *WebService) GetPostDetail(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	postIdInt, _ := strconv.ParseInt(postId, 10, 64)

	result, err := svc.authenticateAndPostClient.GetPostDetail(ctx, &authen_and_post.GetPostRequest{
		PostId: postIdInt,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get post detail FAIL",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func (svc *WebService) UpdatePost(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	postIdInt, _ := strconv.ParseInt(postId, 10, 64)

	var postDataUpdate types.PostRequest

	err := ctx.ShouldBindJSON(&postDataUpdate)
	fmt.Println("postDataUpdate", postDataUpdate)
	if err != nil {
		fmt.Println("ERRORRR", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Data update fail",
			"error":   err.Error(),
		})
		return
	}

	result, err := svc.authenticateAndPostClient.UpdatePost(ctx, &authen_and_post.Post{
		PostId:           postIdInt,
		ContentText:      postDataUpdate.ContextText,
		ContentImagePath: postDataUpdate.ContextImagePath,
		Visible:          postDataUpdate.Visible,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get post detail FAIL",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func (svc *WebService) DeletePost(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	postIdInt, _ := strconv.ParseInt(postId, 10, 64)

	result, err := svc.authenticateAndPostClient.DeletePost(ctx, &authen_and_post.GetPostRequest{
		PostId: postIdInt,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get post detail FAIL",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func (svc *WebService) CommentPost(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	postIdInt, _ := strconv.ParseInt(postId, 10, 64)

	cookie, _ := ctx.Cookie("session_id")
	currentUserId, _ := strconv.ParseInt(cookie, 10, 64)

	var commentPostReq types.CommentPost

	err := ctx.ShouldBindJSON(&commentPostReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Data input fail",
			"error":   err,
		})
		return
	}

	result, err := svc.authenticateAndPostClient.CommentPost(ctx, &authen_and_post.CommentPostRequest{
		PostId:  postIdInt,
		UserId:  currentUserId,
		Content: commentPostReq.Content,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Comment post FAIL",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func (svc *WebService) LikePost(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	postIdInt, err := strconv.ParseInt(postId, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Param post id fail",
			"error":   err,
		})
		return
	}

	cookie, err := ctx.Cookie("session_id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Get session id fail",
			"error":   err,
		})
		return
	}
	currentUserId, _ := strconv.ParseInt(cookie, 10, 64)

	result, err := svc.authenticateAndPostClient.LikePost(ctx, &authen_and_post.LikePostRequest{
		PostId: postIdInt,
		UserId: currentUserId,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Like post FAIL",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Like success",
		"data":    result,
	})
}

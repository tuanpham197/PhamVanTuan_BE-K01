package authen_and_post_svc

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/configs"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/internal/pkg/types"
	"github.com/tuanpham197/PhamVanTuan_BE-K01/pkg/types/proto/pb/authen_and_post"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (a *AuthenticateAndPostService) CheckUserAuthentication(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserResult, error) {
	var user types.User
	err := a.db.First(&user, "user_name = ?", info.UserName).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(info.UserPassword+user.Salt))
	if err != nil {
		return nil, err
	}

	return &authen_and_post.UserResult{
		Status: authen_and_post.UserStatus_OK,
		Info: &authen_and_post.UserDetailInfo{
			UserId:    int64(user.ID),
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
		},
	}, nil

}

func (a *AuthenticateAndPostService) CreateUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error) {
	salt, err := generateRandomString(5)
	if err != nil {
		panic("err")
	}
	passwordHased, errHash := hashPassword(info.Password, salt)
	if errHash != nil {
		panic("Err hash password")
	}
	user := types.User{FirstName: info.FirstName, LastName: info.LastName, UserName: info.UserName, Salt: salt, HashedPassword: passwordHased, DateOfBirth: time.Now()}
	result := a.db.Create(&user)
	if result.Error != nil {
		panic("ERR CREATE " + result.Error.Error())
	}
	return &authen_and_post.UserResult{
		Status: authen_and_post.UserStatus_OK,
		Info: &authen_and_post.UserDetailInfo{
			UserId:    int64(user.ID),
			LastName:  info.LastName,
			FirstName: info.FirstName,
			Email:     info.Email,
		},
	}, nil
}

func (a *AuthenticateAndPostService) EditUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error) {
	fmt.Println("Edit user", info)
	var user types.User
	err := a.db.First(&user, "id = ?", info.UserId).Error
	if err != nil {
		return nil, nil
	}

	if info.FirstName != "" {
		user.FirstName = info.FirstName
	}

	if info.LastName != "" {
		user.LastName = info.LastName
	}

	if info.BirthDay != "" {
		if date, err := time.Parse(time.RFC3339, info.BirthDay); err == nil {
			user.DateOfBirth = date
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	if info.Password != "" {
		if salt, err := generateRandomString(5); err == nil {
			if newPass, err := hashPassword(info.Password, salt); err == nil {
				user.HashedPassword = newPass
			}
		}
	}

	if err := a.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &authen_and_post.UserResult{
		Status: authen_and_post.UserStatus_OK,
		Info: &authen_and_post.UserDetailInfo{
			UserId:    int64(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}, nil
}

func (a *AuthenticateAndPostService) GetUserFollower(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserFollower, error) {
	var user types.User
	var followers []*types.User

	if err := a.db.First(&user, info.UserId).Error; err != nil {
		// Handle the error (user not found)
		return nil, err
	}

	if err := a.db.Model(&user).Association("Follower").Find(&followers); err != nil {
		// Handle the error
		return nil, err
	}

	var followersCustom []*authen_and_post.UserInfo
	for _, follower := range followers {
		followerCustom := &authen_and_post.UserInfo{
			UserId:   int64(follower.ID),
			UserName: follower.UserName,
		}
		followersCustom = append(followersCustom, followerCustom)
	}

	return &authen_and_post.UserFollower{
		Followers: followersCustom,
	}, nil

	// var followers []*UserCustomEntity

	// // Execute a custom SQL query to fetch the followers of the user
	// query := `
	// 	SELECT
	// 		u.id, u.first_name, u.last_name, u.email
	// 	FROM
	// 		users u
	// 	INNER JOIN
	// 		user_follows uf ON uf.friend_id = u.id
	// 	WHERE
	// 		uf.user_id = ?
	// `
	// if err := db.Raw(query, userID).Scan(&followers).Error; err != nil {
	// 	// Handle the error
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve followers"})
	// 	return
	// }
}

func (a *AuthenticateAndPostService) UnFollowUser(ctx context.Context, input *authen_and_post.UserFollowRequest) (*authen_and_post.BoolResp, error) {

	if err := a.db.Where("user_id = ? and friend_id = ?", input.UserId, input.FriendId).Delete(&types.UserFollow{}).Error; err != nil {
		return nil, err
	}

	return &authen_and_post.BoolResp{
		Result: true,
	}, nil
}

func (a *AuthenticateAndPostService) GetPostDetail(ctx context.Context, request *authen_and_post.GetPostRequest) (*authen_and_post.Post, error) {
	var post *authen_and_post.Post

	err := a.db.First(&post).Where(&authen_and_post.Post{PostId: request.PostId}).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (a *AuthenticateAndPostService) FollowUser(ctx context.Context, userFollow *authen_and_post.UserFollowRequest) (*authen_and_post.BoolResp, error) {
	//TODO: validate following exists

	userFollowInput := types.UserFollow{
		UserId:   userFollow.UserId,
		FriendId: userFollow.FriendId,
	}
	result := a.db.Create(&userFollowInput)

	if result.Error != nil {
		return nil, result.Error
	}

	return &authen_and_post.BoolResp{
		Result: true,
	}, nil
}

func (a *AuthenticateAndPostService) CreatePost(ctx context.Context, post *authen_and_post.PostRequest) (*authen_and_post.Post, error) {
	a.db.AutoMigrate(&types.Post{})
	postInput := types.Post{
		UserID:           uint(post.UserId),
		ContentText:      post.ContentText,
		ContentImagePath: post.ContentImagePath,
		Visible:          true,
	}
	result := a.db.Create(&postInput)

	if result.Error != nil {
		panic("ERR CREATE " + result.Error.Error())
	}
	return &authen_and_post.Post{
		PostId:           int64(postInput.ID),
		UserId:           post.UserId,
		ContentText:      post.ContentText,
		ContentImagePath: post.ContentImagePath,
		CreatedAt:        postInput.CreatedAt.String(),
	}, nil
}

func (a *AuthenticateAndPostService) UpdatePost(ctx context.Context, post *authen_and_post.Post) (*authen_and_post.Post, error) {

	var postEntity types.Post
	err := a.db.First(&postEntity).Where("id = ?", post.PostId).Error

	if err != nil {
		return nil, err
	}
	postEntity.ContentText = post.ContentText
	postEntity.ContentImagePath = post.ContentImagePath
	if err := a.db.Save(&postEntity).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (a *AuthenticateAndPostService) DeletePost(ctx context.Context, postRequest *authen_and_post.GetPostRequest) (*authen_and_post.TextResponse, error) {
	err := a.db.Where("id = ?", postRequest.PostId).Delete(&authen_and_post.Post{}).Error

	if err != nil {
		return nil, err
	}

	return &authen_and_post.TextResponse{
		Message: "Delete success",
	}, nil
}

func (a *AuthenticateAndPostService) GetPostFriend(ctx context.Context, userInfo *authen_and_post.UserInfo) (*authen_and_post.PostFriend, error) {
	var result []*authen_and_post.Post
	err := a.db.Where(&types.Post{UserID: uint(userInfo.UserId)}).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &authen_and_post.PostFriend{
		Posts: result,
	}, nil
}

func (a *AuthenticateAndPostService) CommentPost(ctx context.Context, commentRequest *authen_and_post.CommentPostRequest) (*authen_and_post.Post, error) {

	a.db.AutoMigrate(&types.Comment{})
	comment := types.Comment{
		UserID:  uint(commentRequest.UserId),
		Content: commentRequest.Content,
		PostID:  uint(commentRequest.PostId),
	}
	err := a.db.Create(&comment).Error

	if err != nil {
		return nil, err
	}

	var post types.Post

	errGetPost := a.db.First(&post).Where("id = ?", commentRequest.PostId).Error
	if errGetPost != nil {
		return nil, err
	}

	return &authen_and_post.Post{
		PostId:           int64(post.ID),
		UserId:           int64(post.UserID),
		ContentText:      post.ContentText,
		ContentImagePath: post.ContentImagePath,
		Visible:          post.Visible,
		CreatedAt:        post.CreatedAt.String(),
	}, nil
}

func (a *AuthenticateAndPostService) LikePost(ctx context.Context, postRequest *authen_and_post.LikePostRequest) (*authen_and_post.BoolResp, error) {
	a.db.AutoMigrate(&types.Post{})

	// check liked => dislike
	result := a.db.First(&types.Post{}).Where("post_id = ? and user_id = ?", postRequest.PostId, postRequest.UserId)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Fail to check like post")
	}

	//

	likePost := types.Like{
		UserID:    uint(postRequest.UserId),
		PostID:    uint(postRequest.PostId),
		CreatedAt: time.Now(),
	}

	err := a.db.Create(&likePost).Error

	if err != nil {
		return nil, errors.New("Fail to like post")
	}

	// Like success
	return &authen_and_post.BoolResp{
		Result: true,
	}, nil
}

type AuthenticateAndPostService struct {
	authen_and_post.UnimplementedAuthenticateAndPostServer
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthenticateAndPostService(conf *configs.AuthenticateAndPostConfig) (*AuthenticateAndPostService, error) {
	// dsn := "root:123456@tcp(mysql:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(conf.MySQL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("can not connect to db ", err)
		return nil, err
	}
	rd := redis.NewClient(&conf.Redis)
	if rd == nil {
		return nil, fmt.Errorf("can not init redis client")
	}
	return &AuthenticateAndPostService{
		db:    db,
		redis: rd,
	}, nil
}

func generateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)[:length]
	return randomString, nil
}

func hashPassword(password, salt string) (string, error) {
	// Concatenate the password and salt
	passwordWithSalt := []byte(password + salt)

	// Generate the bcrypt hash of the concatenated password and salt
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Lỗi khi mã hóa mật khẩu:", err)
		return "", nil
	}

	return string(hashedPassword), nil
}

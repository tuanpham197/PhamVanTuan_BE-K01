package authen_and_post

import (
	"context"
	"math/rand"

	"github.com/tuanpham197/PhamVanTuan_BE-K01/pkg/types/proto/pb/authen_and_post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type randomClient struct {
	clients []authen_and_post.AuthenticateAndPostClient
}

func (a *randomClient) CheckUserAuthentication(ctx context.Context, in *authen_and_post.UserInfo, opts ...grpc.CallOption) (*authen_and_post.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].CheckUserAuthentication(ctx, in, opts...)
}

func (a *randomClient) CreateUser(ctx context.Context, in *authen_and_post.UserDetailInfo, opts ...grpc.CallOption) (*authen_and_post.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].CreateUser(ctx, in, opts...)
}

func (a *randomClient) EditUser(ctx context.Context, in *authen_and_post.UserDetailInfo, opts ...grpc.CallOption) (*authen_and_post.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].EditUser(ctx, in, opts...)
}

func (a *randomClient) GetUserFollower(ctx context.Context, in *authen_and_post.UserInfo, opts ...grpc.CallOption) (*authen_and_post.UserFollower, error) {
	return a.clients[rand.Intn(len(a.clients))].GetUserFollower(ctx, in, opts...)
}

func (a *randomClient) GetPostDetail(ctx context.Context, in *authen_and_post.GetPostRequest, opts ...grpc.CallOption) (*authen_and_post.Post, error) {
	return a.clients[rand.Intn(len(a.clients))].GetPostDetail(ctx, in, opts...)
}

func (a *randomClient) UnFollowUser(ctx context.Context, in *authen_and_post.UserFollowRequest, opts ...grpc.CallOption) (*authen_and_post.BoolResp, error) {
	return a.clients[rand.Intn(len(a.clients))].UnFollowUser(ctx, in, opts...)
}

func (a *randomClient) FollowUser(ctx context.Context, in *authen_and_post.UserFollowRequest, opts ...grpc.CallOption) (*authen_and_post.BoolResp, error) {
	return a.clients[rand.Intn(len(a.clients))].FollowUser(ctx, in, opts...)
}

func (a *randomClient) CreatePost(ctx context.Context, post *authen_and_post.PostRequest, opts ...grpc.CallOption) (*authen_and_post.Post, error) {
	return a.clients[rand.Intn(len(a.clients))].CreatePost(ctx, post, opts...)
}

func (a *randomClient) DeletePost(ctx context.Context, postRequest *authen_and_post.GetPostRequest, opts ...grpc.CallOption) (*authen_and_post.TextResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].DeletePost(ctx, postRequest, opts...)
}

func (a *randomClient) GetPostFriend(ctx context.Context, userInfo *authen_and_post.UserInfo, opts ...grpc.CallOption) (*authen_and_post.PostFriend, error) {
	return a.clients[rand.Intn(len(a.clients))].GetPostFriend(ctx, userInfo, opts...)
}

func (a *randomClient) UpdatePost(ctx context.Context, post *authen_and_post.Post, opts ...grpc.CallOption) (*authen_and_post.Post, error) {
	return a.clients[rand.Intn(len(a.clients))].UpdatePost(ctx, post, opts...)
}

func (a *randomClient) CommentPost(ctx context.Context, commentPostRequest *authen_and_post.CommentPostRequest, opts ...grpc.CallOption) (*authen_and_post.BoolResp, error) {
	return a.clients[rand.Intn(len(a.clients))].CommentPost(ctx, commentPostRequest, opts...)
}

func (a *randomClient) LikePost(ctx context.Context, likePostRequest *authen_and_post.LikePostRequest, opts ...grpc.CallOption) (*authen_and_post.BoolResp, error) {
	return a.clients[rand.Intn(len(a.clients))].LikePost(ctx, likePostRequest, opts...)
}

func NewClient(hosts []string) (authen_and_post.AuthenticateAndPostClient, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	clients := make([]authen_and_post.AuthenticateAndPostClient, 0, len(hosts))
	for _, host := range hosts {
		conn, err := grpc.Dial(host, opts...)
		if err != nil {
			return nil, err
		}
		client := authen_and_post.NewAuthenticateAndPostClient(conn)
		clients = append(clients, client)
	}
	return &randomClient{clients}, nil
}

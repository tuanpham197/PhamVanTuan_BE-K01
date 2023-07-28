package types

// LoginRequest Login request body
type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Birthday  string `json:"birth_day"`
	Password  string `json:"password"`
}

type PostRequest struct {
	UserId           int64  `json:"user_id"`
	ContextText      string `json:"content_text"`
	ContextImagePath string `json:"content_image_path"`
	Visible          bool   `json:"visible"`
	CreatedAt        string `json:"created_at"`
}

type CommentPost struct {
	UserId  int64  `json:"user_id"`
	Content string `json:"content"`
	PostId  int64  `json:"post_id"`
}

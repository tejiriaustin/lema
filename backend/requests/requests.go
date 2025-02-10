package requests

import "github.com/tejiriaustin/lema/models"

type ()

type (
	CreatePostRequest struct {
		Title  string `json:"title" binding:"required,min=1,max=200"`
		Body   string `json:"body" binding:"required"`
		UserID string `json:"user_id" binding:"required"`
	}

	UpdatePostRequest struct {
		Title string `json:"title" binding:"required,min=1,max=200"`
		Body  string `json:"body" binding:"required"`
	}

	CreateUserRequest struct {
		FullName string         `json:"full_name" binding:"required,min=1,max=200"`
		Email    string         `json:"email" binding:"required,email"`
		Address  models.Address `json:"address" binding:"required"`
	}

	GetUserRequest struct {
		UserID string `json:"user_id" binding:"required"`
	}
	GetPostsRequest struct {
	}
)

type ()

type ()

package service

import (
	"context"

	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
)

type (
	UserServiceInterface interface {
		CreateUser(ctx context.Context,
			input CreateUserInput,
			userRepo repository.RepoInterface[models.User],
		) (*models.User, error)

		GetUsers(ctx context.Context,
			input GetUsersInput,
			userRepo repository.RepoInterface[models.User],
		) ([]*models.User, *repository.Paginator, error)

		GetUserByID(ctx context.Context,
			id string,
			userRepo repository.RepoInterface[models.User],
		) (*models.User, error)

		GetUserCount(ctx context.Context,
			userRepo repository.RepoInterface[models.User],
		) (int64, error)
	}

	PostServiceInterface interface {
		CreatePost(ctx context.Context,
			input CreatePostInput,
			postRepo repository.RepoInterface[models.Post],
		) (*models.Post, error)

		GetUserPosts(ctx context.Context,
			input GetUserPostInput,
			postRepo repository.RepoInterface[models.Post],
		) ([]*models.Post, *repository.Paginator, error)

		DeletePost(ctx context.Context,
			userID string,
			postRepo repository.RepoInterface[models.Post],
		) error
	}
)

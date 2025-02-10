package service

import (
	"context"
	"github.com/tejiriaustin/lema/logger"
	"go.uber.org/zap"
	"strconv"

	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
)

type (
	PostService struct {
		_          struct{}
		lemaLogger logger.Logger
	}

	CreatePostInput struct {
		Title  string
		Body   string
		UserID string
	}
	GetUserPostInput struct {
		Pager
		UserID string
	}
)

var _ PostServiceInterface = (*PostService)(nil)

func NewPostService(lemaLogger logger.Logger) PostServiceInterface {
	return &PostService{
		lemaLogger: lemaLogger,
	}
}

func (s *PostService) CreatePost(ctx context.Context,
	input CreatePostInput,
	postRepo repository.RepoInterface[models.Post],
) (*models.Post, error) {

	post := models.Post{
		UserID: input.UserID,
		Title:  input.Title,
		Body:   input.Body,
	}

	createdPost, err := postRepo.Create(ctx, post)
	if err != nil {
		s.lemaLogger.Error("failed to create user",
			err,
			zap.String("user_id", input.UserID),
			zap.String("title", input.Title),
			zap.String("body_length", strconv.Itoa(len(input.Body))),
		)
		return nil, err
	}
	return createdPost, nil
}

func (s *PostService) GetUserPosts(ctx context.Context,
	input GetUserPostInput,
	postRepo repository.RepoInterface[models.Post],
) ([]*models.Post, *repository.Paginator, error) {
	filter := repository.NewQueryFilter().Where("user_id = ?", input.UserID)

	posts, paginate, err := postRepo.FindManyPaginated(ctx, filter, input.Page, input.PerPage)
	if err != nil {
		s.lemaLogger.Error("failed to create user", err, zap.String("user_id", input.UserID))
		return nil, nil, err
	}
	return posts, paginate, nil
}

func (s *PostService) DeletePost(ctx context.Context,
	postID string,
	postRepo repository.RepoInterface[models.Post],
) error {
	filter := repository.NewQueryFilter().Where("id = ?", postID)

	err := postRepo.DeleteMany(ctx, filter)
	if err != nil {
		s.lemaLogger.Error("failed to create user", err, zap.String("post_id", postID))
		return err
	}
	return nil
}

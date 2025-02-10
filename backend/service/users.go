package service

import (
	"context"
	"errors"
	"github.com/tejiriaustin/lema/logger"
	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
	"go.uber.org/zap"
)

type (
	UserService struct {
		_          struct{}
		lemaLogger logger.Logger
	}

	CreateUserInput struct {
		FullName string
		Email    string
		Address  *models.Address
	}

	GetUsersInput struct {
		Pager
		Filters GetUsersFilters
	}

	GetUsersFilters struct{}
)

var _ UserServiceInterface = (*UserService)(nil)

func NewUserService(lemaLogger logger.Logger) UserServiceInterface {
	return &UserService{
		lemaLogger: lemaLogger,
	}
}

func (s *UserService) CreateUser(ctx context.Context,
	input CreateUserInput,
	userRepo repository.RepoInterface[models.User],
) (*models.User, error) {

	user := models.User{
		FullName: input.FullName,
		Email:    input.Email,
		Address:  input.Address,
	}

	createdUser, err := userRepo.Create(ctx, user)
	if err != nil {
		s.lemaLogger.Error("failed to create post",
			err,
			zap.String("full_name", input.FullName),
			zap.String("email", input.Email),
			zap.String("address", input.Address.String()),
		)
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) GetUsers(ctx context.Context,
	input GetUsersInput,
	userRepo repository.RepoInterface[models.User],
) ([]*models.User, *repository.Paginator, error) {

	users, paginate, err := userRepo.FindManyPaginated(ctx, nil, input.Page, input.PerPage, "Address")
	if err != nil {
		s.lemaLogger.Error("failed to get users", err)
		return nil, nil, err
	}

	return users, paginate, nil
}

func (s *UserService) GetUserByID(ctx context.Context,
	userID string,
	userRepo repository.RepoInterface[models.User],
) (*models.User, error) {
	filter := repository.NewQueryFilter().Where("id = ?", userID)

	user, err := userRepo.FindOne(ctx, filter, "Address")
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *UserService) GetUserCount(ctx context.Context,
	userRepo repository.RepoInterface[models.User],
) (int64, error) {

	filter := repository.NewQueryFilter()

	return userRepo.Count(ctx, filter)
}

package service

import (
	"context"
	"errors"
	"github.com/tejiriaustin/lema/logger"
	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
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
		Name:    input.FullName,
		Email:   input.Email,
		Address: input.Address,
	}

	filter := repository.NewQueryFilter().Where("email = ?", user.Email)

	foundUser, err := userRepo.FindOne(ctx, filter)
	if foundUser != nil {
		s.lemaLogger.Error("found user with matching email",
			logger.WithField("email", input.Email),
		)
		return nil, errors.New("A user with this email already exists")
	}

	createdUser, err := userRepo.Create(ctx, user)
	if err != nil {
		s.lemaLogger.Error("failed to create USER",
			logger.WithField("err", err),
			logger.WithField("full_name", input.FullName),
			logger.WithField("email", input.Email),
			logger.WithField("address", input.Address.String()),
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
		s.lemaLogger.Error("failed to get users", logger.WithField("err", err))
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
		s.lemaLogger.Error("failed to get user by id", logger.WithField("err", err))
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

package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/tejiriaustin/lema/logger"
	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
	"github.com/tejiriaustin/lema/service"
	"github.com/tejiriaustin/lema/testutils"
	loggermocks "github.com/tejiriaustin/lema/testutils/mocks/logger"
	repomocks "github.com/tejiriaustin/lema/testutils/mocks/repository"
)

type UserServiceTestSuite struct {
	testutils.BaseSuite
	service service.UserServiceInterface
}

func TestUserService(t *testing.T) {
	mockLogger := new(loggermocks.Logger)
	testService := &UserServiceTestSuite{
		service: service.NewUserService(mockLogger),
	}
	suite.Run(t, testService)
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name        string
			input       func() service.CreateUserInput
			output      func() *models.User
			expectError bool
			setupMock   func(*repomocks.RepoInterface[models.User], *loggermocks.Logger)
		}

		address := &models.Address{
			Street: "123 Main St",
			City:   "Example City",
			State:  "EX",
		}

		testCases := []testCase{
			{
				name: "successfully create a user",
				input: func() service.CreateUserInput {
					return service.CreateUserInput{
						FullName: "John Doe",
						Email:    "john@example.com",
						Address:  address,
					}
				},
				output: func() *models.User {
					return &models.User{
						FullName: "John Doe",
						Email:    "john@example.com",
						Address:  address,
					}
				},
				expectError: false,
				setupMock: func(repo *repomocks.RepoInterface[models.User], mockLogger *loggermocks.Logger) {
					repo.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)

					repo.On("Create", mock.Anything, mock.MatchedBy(func(u models.User) bool {
						return u.FullName == "John Doe" && u.Email == "john@example.com"
					})).Return(&models.User{
						FullName: "John Doe",
						Email:    "john@example.com",
						Address:  address,
					}, nil)
				},
			},
			{
				name: "user with email already exists",
				input: func() service.CreateUserInput {
					return service.CreateUserInput{
						FullName: "John Doe",
						Email:    "john@example.com",
						Address:  address,
					}
				},
				output: func() *models.User {
					return nil
				},
				expectError: true,
				setupMock: func(repo *repomocks.RepoInterface[models.User], mockLogger *loggermocks.Logger) {
					mockLogger.On("Error",
						"found user with matching email",
						logger.Field{Key: "email", Value: "john@example.com"},
					).Return()

					repo.On("FindOne", mock.Anything, mock.Anything).Return(&models.User{
						FullName: "John Doe",
						Email:    "john@example.com",
						Address:  address,
					}, nil)
				},
			},
			{
				name: "error creating user",
				input: func() service.CreateUserInput {
					return service.CreateUserInput{
						FullName: "John Doe",
						Email:    "john@example.com",
						Address:  address,
					}
				},
				output: func() *models.User {
					return nil
				},
				expectError: true,
				setupMock: func(repo *repomocks.RepoInterface[models.User], mockLogger *loggermocks.Logger) {
					repo.On("FindOne", mock.Anything, mock.Anything).Return(nil, nil)
					mockLogger.On("Error",
						"failed to create USER",
						logger.Field{Key: "err", Value: errors.New("database error")},
						logger.Field{Key: "full_name", Value: "John Doe"},
						logger.Field{Key: "email", Value: "john@example.com"},
						logger.Field{Key: "address", Value: address.String()},
					).Return()

					repo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("database error"))
				},
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				mockLogger := new(loggermocks.Logger)
				userRepo := new(repomocks.RepoInterface[models.User])

				svc := service.NewUserService(mockLogger)
				input := tc.input()
				tc.setupMock(userRepo, mockLogger)

				user, err := svc.CreateUser(ctx, input, userRepo)

				if tc.expectError {
					suite.NotNil(err)
					suite.Nil(user)
					return
				}
				suite.Nil(err)
				suite.Equal(tc.output(), user)

				userRepo.AssertExpectations(suite.T())
				mockLogger.AssertExpectations(suite.T())
			})
		}
	})
}

func (suite *UserServiceTestSuite) TestGetUsers() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name        string
			input       func() service.GetUsersInput
			output      func() []*models.User
			paginate    func() *repository.Paginator
			setupMock   func(*repomocks.RepoInterface[models.User])
			expectError bool
		}

		testCases := []testCase{
			{
				name: "successfully get users",
				input: func() service.GetUsersInput {
					return service.GetUsersInput{
						Pager: service.Pager{
							Page:    1,
							PerPage: 10,
						},
					}
				},
				output: func() []*models.User {
					return []*models.User{
						{
							FullName: "John Doe",
							Email:    "john@example.com",
						},
						{
							FullName: "Jane Doe",
							Email:    "jane@example.com",
						},
					}
				},
				paginate: func() *repository.Paginator {
					return &repository.Paginator{
						CurrentPage: 1,
						PerPage:     10,
						TotalRows:   2,
					}
				},
				setupMock: func(repo *repomocks.RepoInterface[models.User]) {
					repo.On("FindManyPaginated",
						mock.Anything,
						mock.Anything,
						int64(1),
						int64(10),
						"Address",
					).Return([]*models.User{
						{
							FullName: "John Doe",
							Email:    "john@example.com",
						},
						{
							FullName: "Jane Doe",
							Email:    "jane@example.com",
						},
					}, &repository.Paginator{
						CurrentPage: 1,
						PerPage:     10,
						TotalRows:   2,
					}, nil)
				},
				expectError: false,
			},
			{
				name: "error getting users",
				input: func() service.GetUsersInput {
					return service.GetUsersInput{
						Pager: service.Pager{
							Page:    1,
							PerPage: 10,
						},
					}
				},
				output: func() []*models.User {
					return nil
				},
				paginate: func() *repository.Paginator {
					return nil
				},
				setupMock: func(repo *repomocks.RepoInterface[models.User]) {
					repo.On("FindManyPaginated",
						mock.Anything,
						mock.Anything,
						int64(1),
						int64(10),
						"Address",
					).Return(nil, nil, errors.New("database error"))
				},
				expectError: true,
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				mockLogger := new(loggermocks.Logger)
				userRepo := new(repomocks.RepoInterface[models.User])

				if tc.expectError {
					mockLogger.On("Error",
						"failed to get users",
						logger.Field{Key: "err", Value: errors.New("database error")},
					).Return()
				}

				svc := service.NewUserService(mockLogger)
				input := tc.input()
				tc.setupMock(userRepo)

				users, paginator, err := svc.GetUsers(ctx, input, userRepo)

				if tc.expectError {
					suite.NotNil(err)
					suite.Nil(users)
					suite.Nil(paginator)
				} else {
					suite.Nil(err)
					suite.Equal(tc.output(), users)
					suite.Equal(tc.paginate(), paginator)
				}

				userRepo.AssertExpectations(suite.T())
				mockLogger.AssertExpectations(suite.T())
			})
		}
	})
}

func (suite *UserServiceTestSuite) TestGetUserByID() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name        string
			userID      string
			output      func() *models.User
			setupMock   func(*repomocks.RepoInterface[models.User])
			expectError bool
		}

		testCases := []testCase{
			{
				name:   "successfully get user by ID",
				userID: "user-123",
				output: func() *models.User {
					return &models.User{
						FullName: "John Doe",
						Email:    "john@example.com",
					}
				},
				setupMock: func(repo *repomocks.RepoInterface[models.User]) {
					repo.On("FindOne",
						mock.Anything,
						mock.MatchedBy(func(q *repository.Query) bool {
							return true // Add more specific matching if needed
						}),
						"Address",
					).Return(&models.User{
						FullName: "John Doe",
						Email:    "john@example.com",
					}, nil)
				},
				expectError: false,
			},
			{
				name:   "user not found",
				userID: "non-existent",
				output: func() *models.User {
					return nil
				},
				setupMock: func(repo *repomocks.RepoInterface[models.User]) {
					repo.On("FindOne",
						mock.Anything,
						mock.Anything,
						"Address",
					).Return(nil, errors.New("not found"))
				},
				expectError: true,
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				mockLogger := new(loggermocks.Logger)
				userRepo := new(repomocks.RepoInterface[models.User])

				if tc.expectError {
					mockLogger.On("Error",
						"failed to get user by id",
						logger.Field{Key: "err", Value: errors.New("not found")},
					).Return()
				}

				svc := service.NewUserService(mockLogger)
				tc.setupMock(userRepo)

				user, err := svc.GetUserByID(ctx, tc.userID, userRepo)

				if tc.expectError {
					suite.NotNil(err)
					suite.Nil(user)
					suite.Equal("user not found", err.Error())
				} else {
					suite.Nil(err)
					suite.Equal(tc.output(), user)
				}

				userRepo.AssertExpectations(suite.T())
				mockLogger.AssertExpectations(suite.T())
			})
		}
	})
}

func (suite *UserServiceTestSuite) TestGetUserCount() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name        string
			output      int64
			setupMock   func(*repomocks.RepoInterface[models.User])
			expectError bool
		}

		testCases := []testCase{
			{
				name:   "successfully get user count",
				output: 42,
				setupMock: func(repo *repomocks.RepoInterface[models.User]) {
					repo.On("Count",
						mock.Anything,
						mock.MatchedBy(func(q *repository.Query) bool {
							return true
						}),
					).Return(int64(42), nil)
				},
				expectError: false,
			},
			{
				name:   "error getting user count",
				output: 0,
				setupMock: func(repo *repomocks.RepoInterface[models.User]) {
					repo.On("Count",
						mock.Anything,
						mock.Anything,
					).Return(int64(0), errors.New("database error"))
				},
				expectError: true,
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				userRepo := new(repomocks.RepoInterface[models.User])
				tc.setupMock(userRepo)

				count, err := suite.service.GetUserCount(ctx, userRepo)

				if tc.expectError {
					suite.NotNil(err)
					suite.Equal(int64(0), count)
				} else {
					suite.Nil(err)
					suite.Equal(tc.output, count)
				}

				userRepo.AssertExpectations(suite.T())
			})
		}
	})
}

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

type PostServiceTestSuite struct {
	testutils.BaseSuite
	service service.PostServiceInterface
}

func TestPostService(t *testing.T) {
	mockLogger := new(loggermocks.Logger)

	testService := &PostServiceTestSuite{
		service: service.NewPostService(mockLogger),
	}
	suite.Run(t, testService)
}

func (suite *PostServiceTestSuite) TestCreatePost() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name        string
			input       func() service.CreatePostInput
			output      func() models.Post
			expectError bool
			setupMock   func(*repomocks.RepoInterface[models.Post])
		}

		testCases := []testCase{
			{
				name: "successfully create a post",
				input: func() service.CreatePostInput {
					return service.CreatePostInput{
						Title:  "I Got a Letter",
						Body:   "Lorem ipsum dolor sit amet. ",
						UserID: "05c-90df-4f23-999a-28469f2a58e3",
					}
				},
				output: func() models.Post {
					return models.Post{
						UserID: "05c-90df-4f23-999a-28469f2a58e3",
						Title:  "I Got a Letter",
						Body:   "Lorem ipsum dolor sit amet. ",
					}
				},
				expectError: false,
				setupMock: func(repo *repomocks.RepoInterface[models.Post]) {
					repo.On("Create", mock.Anything, mock.MatchedBy(func(i models.Post) bool {
						return i.Title == "I Got a Letter" && i.Body == "Lorem ipsum dolor sit amet. "
					})).Return(&models.Post{
						UserID: "05c-90df-4f23-999a-28469f2a58e3",
						Title:  "I Got a Letter",
						Body:   "Lorem ipsum dolor sit amet. ",
					}, nil)
				},
			},
			{
				name: "error creating post",
				input: func() service.CreatePostInput {
					return service.CreatePostInput{
						Title:  "I Got a Letter",
						Body:   "Lorem ipsum dolor sit amet. ",
						UserID: "05c-90df-4f23-999a-28469f2a58e3",
					}
				},
				output: func() models.Post {
					return models.Post{}
				},
				expectError: true,
				setupMock: func(repo *repomocks.RepoInterface[models.Post]) {
					repo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("database error"))
				},
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				mockLogger := new(loggermocks.Logger)
				postRepo := new(repomocks.RepoInterface[models.Post])

				if tc.expectError {
					mockLogger.On("Error",
						"failed to create post",
						logger.Field{Key: "err", Value: errors.New("database error")},
						logger.Field{Key: "user_id", Value: "05c-90df-4f23-999a-28469f2a58e3"},
						logger.Field{Key: "title", Value: "I Got a Letter"},
						logger.Field{Key: "body_length", Value: "28"},
					).Return()
				}

				svc := service.NewPostService(mockLogger)

				input := tc.input()
				tc.setupMock(postRepo)

				_, err := svc.CreatePost(ctx, input, postRepo)

				if tc.expectError {
					suite.NotNil(err)
				} else {
					suite.Nil(err)
				}

				postRepo.AssertExpectations(suite.T())
				mockLogger.AssertExpectations(suite.T())
			})
		}
	})
}

func (suite *PostServiceTestSuite) TestGetUserPosts() {
	suite.NotPanics(func() {
		ctx := context.Background()

		type testCase struct {
			name        string
			input       func() service.GetUserPostInput
			output      func() []*models.Post
			paginate    func() *repository.Paginator
			setupMock   func(*repomocks.RepoInterface[models.Post])
			expectError bool
		}

		testCases := []testCase{
			{
				name: "successfully get user posts",
				input: func() service.GetUserPostInput {
					return service.GetUserPostInput{
						UserID: "05c-90df-4f23-999a-28469f2a58e3",
						Pager: service.Pager{
							Page:    1,
							PerPage: 10,
						},
					}
				},
				output: func() []*models.Post {
					return []*models.Post{
						{
							UserID: "05c-90df-4f23-999a-28469f2a58e3",
							Title:  "First Post",
							Body:   "First post content",
						},
						{
							UserID: "05c-90df-4f23-999a-28469f2a58e3",
							Title:  "Second Post",
							Body:   "Second post content",
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
				setupMock: func(repo *repomocks.RepoInterface[models.Post]) {
					repo.On("FindManyPaginated",
						mock.Anything,
						mock.MatchedBy(func(f *repository.Query) bool {
							return true
						}),
						int64(1),
						int64(10),
					).Return([]*models.Post{
						{
							UserID: "05c-90df-4f23-999a-28469f2a58e3",
							Title:  "First Post",
							Body:   "First post content",
						},
						{
							UserID: "05c-90df-4f23-999a-28469f2a58e3",
							Title:  "Second Post",
							Body:   "Second post content",
						},
					}, &repository.Paginator{
						CurrentPage: 1,
						PerPage:     10,
						TotalRows:   2,
					}, nil)
				},
				expectError: false,
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				input := tc.input()
				expectedOutput := tc.output()
				expectedPaginate := tc.paginate()

				postRepo := new(repomocks.RepoInterface[models.Post])
				tc.setupMock(postRepo)

				posts, paginator, err := suite.service.GetUserPosts(ctx, input, postRepo)

				if tc.expectError {
					suite.NotNil(err)
					suite.Nil(posts)
					suite.Nil(paginator)
					return
				}
				suite.Nil(err)
				suite.Equal(expectedOutput, posts)
				suite.Equal(expectedPaginate, paginator)

				postRepo.AssertExpectations(suite.T())
			})
		}
	})
}

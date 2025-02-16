package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/tejiriaustin/lema/controllers"
	"github.com/tejiriaustin/lema/env"
	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
	"github.com/tejiriaustin/lema/requests"
	"github.com/tejiriaustin/lema/service"
	"github.com/tejiriaustin/lema/testutils"
	servicemocks "github.com/tejiriaustin/lema/testutils/mocks/service"
)

type PostControllerTestSuite struct {
	testutils.BaseSuite
	controller *controllers.PostController
	conf       *env.Environment
}

func TestPostController(t *testing.T) {
	conf := &env.Environment{}
	suite.Run(t, &PostControllerTestSuite{
		BaseSuite:  testutils.BaseSuite{},
		controller: controllers.NewPostController(conf),
		conf:       conf,
	})
}

func (suite *PostControllerTestSuite) setupTest() (*gin.Engine, *servicemocks.UserServiceInterface, *servicemocks.PostServiceInterface, *repository.Repository[models.User], *repository.Repository[models.Post]) {
	mockUserSvc := new(servicemocks.UserServiceInterface)
	mockPostSvc := new(servicemocks.PostServiceInterface)
	userRepo := &repository.Repository[models.User]{}
	postsRepo := &repository.Repository[models.Post]{}

	gin.SetMode(gin.TestMode)
	router := gin.New()

	return router, mockUserSvc, mockPostSvc, userRepo, postsRepo
}

func (suite *PostControllerTestSuite) TestCreatePost() {
	suite.NotPanics(func() {
		type testCase struct {
			name         string
			input        requests.CreatePostRequest
			setupMocks   func(*servicemocks.UserServiceInterface, *servicemocks.PostServiceInterface)
			expectedCode int
			expectedMsg  string
		}

		testCases := []testCase{
			{
				name: "successfully create post",
				input: requests.CreatePostRequest{
					Title:  "Test Post",
					Body:   "Test Body",
					UserID: "user123",
				},
				setupMocks: func(userSvc *servicemocks.UserServiceInterface, postSvc *servicemocks.PostServiceInterface) {
					userSvc.On("GetUserByID",
						mock.Anything,
						"user123",
						mock.Anything,
					).Return(&models.User{
						Name: "Test User",
					}, nil)

					postSvc.On("CreatePost",
						mock.Anything,
						service.CreatePostInput{
							Title:  "Test Post",
							Body:   "Test Body",
							UserID: "user123",
						},
						mock.Anything,
					).Return(&models.Post{
						Title:  "Test Post",
						Body:   "Test Body",
						UserID: "user123",
					}, nil)
				},
				expectedCode: http.StatusOK,
				expectedMsg:  "successful",
			},
			{
				name: "invalid user ID",
				input: requests.CreatePostRequest{
					Title:  "Test Post",
					Body:   "Test Body",
					UserID: "invalid_user",
				},
				setupMocks: func(userSvc *servicemocks.UserServiceInterface, postSvc *servicemocks.PostServiceInterface) {
					userSvc.On("GetUserByID",
						mock.Anything,
						"invalid_user",
						mock.Anything,
					).Return(nil, errors.New("user not found"))
				},
				expectedCode: http.StatusBadRequest,
				expectedMsg:  "Invalid User ID",
			},
			{
				name: "error creating post",
				input: requests.CreatePostRequest{
					Title:  "Test Post",
					Body:   "Test Body",
					UserID: "user123",
				},
				setupMocks: func(userSvc *servicemocks.UserServiceInterface, postSvc *servicemocks.PostServiceInterface) {
					userSvc.On("GetUserByID",
						mock.Anything,
						"user123",
						mock.Anything,
					).Return(&models.User{
						Name: "Test User",
					}, nil)

					postSvc.On("CreatePost",
						mock.Anything,
						service.CreatePostInput{
							Title:  "Test Post",
							Body:   "Test Body",
							UserID: "user123",
						},
						mock.Anything,
					).Return(nil, errors.New("failed to create post"))
				},
				expectedCode: http.StatusInternalServerError,
				expectedMsg:  "failed to create post",
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				// Setup fresh instances for each test case
				router, mockUserSvc, mockPostSvc, userRepo, postsRepo := suite.setupTest()

				// Setup the route for this test case
				router.POST("/posts", suite.controller.CreatePost(
					mockUserSvc,
					mockPostSvc,
					userRepo,
					postsRepo,
				))

				// Setup mocks
				tc.setupMocks(mockUserSvc, mockPostSvc)

				// Create request
				body, _ := json.Marshal(tc.input)
				req, _ := http.NewRequestWithContext(
					context.Background(),
					"POST",
					"/posts",
					bytes.NewBuffer(body),
				)
				req.Header.Set("Content-Type", "application/json")

				// Perform request
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				// Assert response
				suite.Equal(tc.expectedCode, w.Code)

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				suite.NoError(err)
				suite.Equal(tc.expectedMsg, response["message"])

				// Verify mocks
				mockUserSvc.AssertExpectations(suite.T())
				mockPostSvc.AssertExpectations(suite.T())
			})
		}
	})
}

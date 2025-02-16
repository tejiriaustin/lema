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

type UserControllerTestSuite struct {
	testutils.BaseSuite
	controller *controllers.UserController
	conf       *env.Environment
}

func TestUserController(t *testing.T) {
	conf := &env.Environment{}
	suite.Run(t, &UserControllerTestSuite{
		BaseSuite:  testutils.BaseSuite{},
		controller: controllers.NewUserController(conf),
		conf:       conf,
	})
}

func (suite *UserControllerTestSuite) setupTest() (*gin.Engine, *servicemocks.UserServiceInterface, *repository.Repository[models.User]) {
	mockUserSvc := new(servicemocks.UserServiceInterface)
	UsersRepo := &repository.Repository[models.User]{}

	gin.SetMode(gin.TestMode)
	router := gin.New()

	return router, mockUserSvc, UsersRepo
}

func (suite *UserControllerTestSuite) TestCreateUser() {
	suite.NotPanics(func() {
		type testCase struct {
			name         string
			input        requests.CreateUserRequest
			setupMocks   func(*servicemocks.UserServiceInterface)
			expectedCode int
			expectedMsg  string
		}

		address := models.Address{
			Street:  "Street",
			City:    "City",
			State:   "State",
			ZipCode: "ZipCode",
		}

		testCases := []testCase{
			{
				name: "successfully create user",
				input: requests.CreateUserRequest{
					FullName: "Test User",
					Email:    "test@example.com",
					Address:  address,
				},
				setupMocks: func(userSvc *servicemocks.UserServiceInterface) {
					userSvc.On("CreateUser",
						mock.Anything,
						service.CreateUserInput{
							FullName: "Test User",
							Email:    "test@example.com",
							Address:  &address,
						},
						mock.Anything,
					).Return(&models.User{
						Name:    "Test User",
						Email:   "test@example.com",
						Address: &address,
					}, nil)
				},
				expectedCode: http.StatusOK,
				expectedMsg:  "successful",
			},
			{
				name: "user already exists",
				input: requests.CreateUserRequest{
					FullName: "Test User",
					Email:    "existing@example.com",
					Address:  address,
				},
				setupMocks: func(userSvc *servicemocks.UserServiceInterface) {
					userSvc.On("CreateUser",
						mock.Anything,
						service.CreateUserInput{
							FullName: "Test User",
							Email:    "existing@example.com",
							Address:  &address,
						},
						mock.Anything,
					).Return(nil, errors.New("A user with this email already exists"))
				},
				expectedCode: http.StatusInternalServerError,
				expectedMsg:  "A user with this email already exists",
			},
			{
				name: "internal server error",
				input: requests.CreateUserRequest{
					FullName: "Test User",
					Email:    "test@example.com",
					Address:  address,
				},
				setupMocks: func(userSvc *servicemocks.UserServiceInterface) {
					userSvc.On("CreateUser",
						mock.Anything,
						service.CreateUserInput{
							FullName: "Test User",
							Email:    "test@example.com",
							Address:  &address,
						},
						mock.Anything,
					).Return(nil, errors.New("failed to create user"))
				},
				expectedCode: http.StatusInternalServerError,
				expectedMsg:  "failed to create user",
			},
		}

		for _, tc := range testCases {
			suite.Run(tc.name, func() {
				router, mockUserSvc, usersRepo := suite.setupTest()

				router.POST("/users", suite.controller.CreateUser(
					mockUserSvc,
					usersRepo,
				))

				tc.setupMocks(mockUserSvc)

				body, _ := json.Marshal(tc.input)
				req, _ := http.NewRequestWithContext(
					context.Background(),
					http.MethodPost,
					"/users",
					bytes.NewBuffer(body),
				)
				req.Header.Set("Content-Type", "application/json")

				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)

				suite.Equal(tc.expectedCode, w.Code)

				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				suite.NoError(err)
				suite.Equal(tc.expectedMsg, response["message"])

				mockUserSvc.AssertExpectations(suite.T())
			})
		}
	})
}

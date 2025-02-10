package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/lema/env"
	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
	"github.com/tejiriaustin/lema/requests"
	"github.com/tejiriaustin/lema/response"
	"github.com/tejiriaustin/lema/service"
)

type UserController struct {
	conf *env.Environment
}

func NewUserController(conf *env.Environment) *UserController {
	return &UserController{
		conf: conf,
	}
}

func (c *UserController) CreateUser(
	userService service.UserServiceInterface,
	usersRepo *repository.Repository[models.User],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		input := service.CreateUserInput{
			FullName: req.FullName,
			Email:    req.Email,
			Address:  &req.Address,
		}

		user, err := userService.CreateUser(ctx, input, usersRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleUserResponse(user))
	}
}

func (c *UserController) GetUser(
	userService service.UserServiceInterface,
	usersRepo *repository.Repository[models.User],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.GetUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		user, err := userService.GetUserByID(ctx, req.UserID, usersRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleUserResponse(user))
	}
}

func (c *UserController) GetUsers(
	userService service.UserServiceInterface,
	usersRepo *repository.Repository[models.User],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := service.GetUsersInput{
			Pager: service.Pager{
				Page:    service.GetPageNumberFromContext(ctx),
				PerPage: service.GetPageSizeLimitFromContext(ctx),
			},
		}

		users, paginate, err := userService.GetUsers(ctx, input, usersRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		payload := map[string]interface{}{
			"paginationData": paginate,
			"users":          response.MultipleUserResponse(users),
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", payload)
	}
}

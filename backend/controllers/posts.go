package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/tejiriaustin/lema/env"
	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
	"github.com/tejiriaustin/lema/requests"
	"github.com/tejiriaustin/lema/response"
	"github.com/tejiriaustin/lema/service"
)

type PostController struct {
	conf *env.Environment
}

func NewPostController(conf *env.Environment) *PostController {
	return &PostController{
		conf: conf,
	}
}

func (c *PostController) CreatePost(
	userService service.UserServiceInterface,
	postService service.PostServiceInterface,
	userRepo *repository.Repository[models.User],
	postsRepo *repository.Repository[models.Post],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.CreatePostRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		user, err := userService.GetUserByID(ctx, req.UserID, userRepo)
		if err != nil || user == nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Invalid User ID", nil)
			return
		}

		input := service.CreatePostInput{
			Title:  req.Title,
			Body:   req.Body,
			UserID: req.UserID,
		}

		post, err := postService.CreatePost(ctx, input, postsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SinglePostResponse(post))
	}
}

func (c *PostController) GetPosts(
	userService service.UserServiceInterface,
	postService service.PostServiceInterface,
	userRepo *repository.Repository[models.User],
	postsRepo *repository.Repository[models.Post],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Query("user_id")
		if userID == "" {
			response.FormatResponse(ctx, http.StatusBadRequest, "user id is required", nil)
			return
		}

		user, err := userService.GetUserByID(ctx, userID, userRepo)
		if err != nil || user == nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Invalid User ID", nil)
			return
		}

		input := service.GetUserPostInput{
			UserID: userID,
			Pager: service.Pager{
				Page:    service.GetPageNumberFromContext(ctx),
				PerPage: service.GetPageSizeLimitFromContext(ctx),
			},
		}

		posts, paginationData, err := postService.GetUserPosts(ctx, input, postsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		payload := map[string]interface{}{
			"paginationData": paginationData,
			"posts":          response.MultiplePostResponse(posts),
			"user":           response.SingleUserResponse(user),
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", payload)
	}
}

func (c *PostController) DeletePost(
	postService service.PostServiceInterface,
	postsRepo repository.RepoInterface[models.Post],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postID := ctx.Param("id")
		if postID == "" {
			response.FormatResponse(ctx, http.StatusBadRequest, "post id is required", nil)
			return
		}

		err := postService.DeletePost(ctx.Request.Context(), postID, postsRepo)
		if err != nil {
			switch {
			case errors.Is(err, repository.ErrNotFound):
				response.FormatResponse(ctx, http.StatusNotFound, "post not found", nil)
			default:
				response.FormatResponse(ctx, http.StatusInternalServerError, "failed to delete post", nil)
			}
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "post deleted successfully", nil)
	}
}

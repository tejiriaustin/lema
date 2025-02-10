package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/lema/env"
	"github.com/tejiriaustin/lema/repository"
	"github.com/tejiriaustin/lema/response"
	"github.com/tejiriaustin/lema/service"
)

func BindRoutes(
	ctx context.Context,
	routerEngine *gin.Engine,
	sc *service.Container, // sc stands for Service Container
	repo *repository.Container,
	conf *env.Environment,
) {

	controllers := New(ctx, conf)

	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		response.FormatResponse(c, http.StatusOK, "OK", nil)
	})

	users := r.Group("/users")
	{
		users.POST("", controllers.UserController.CreateUser(sc.UserService, repo.UserRepo)) // POST /api/v1
		users.GET("", controllers.UserController.GetUsers(sc.UserService, repo.UserRepo))    // GET /api/v1/users?pageNumber=0&pageSize=10
	}

	posts := r.Group("/posts")
	{
		posts.POST("", controllers.PostController.CreatePost(sc.UserService, sc.PostService, repo.UserRepo, repo.PostRepo)) // POST /api/v1/posts
		posts.GET("", controllers.PostController.GetPosts(sc.UserService, sc.PostService, repo.UserRepo, repo.PostRepo))    // GET /api/v1/posts?userId=1
		posts.DELETE("/:id", controllers.PostController.DeletePost(sc.PostService, repo.PostRepo))                          // DELETE /api/v1/posts/:id
	}
}

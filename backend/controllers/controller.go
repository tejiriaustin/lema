package controllers

import (
	"context"
	"github.com/tejiriaustin/lema/env"
)

type (
	Controller struct {
		conf           *env.Environment
		UserController *UserController
		PostController *PostController
	}
)

func New(ctx context.Context, conf *env.Environment) *Controller {
	return &Controller{
		UserController: NewUserController(conf),
		PostController: NewPostController(conf),
	}
}

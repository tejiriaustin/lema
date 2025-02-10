package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/lema/controllers"
	"github.com/tejiriaustin/lema/env"
	"github.com/tejiriaustin/lema/middleware"
	"github.com/tejiriaustin/lema/repository"
	"github.com/tejiriaustin/lema/service"
)

func Start(
	ctx context.Context,
	service *service.Container,
	repo *repository.Container,
	conf *env.Environment,
) {
	router := gin.New()

	router.Use(
		middleware.CORSMiddleware(),
		middleware.DefaultStructuredLogs(),
		middleware.ReadPaginationOptions(),
	)
	log.Println("starting server...")

	controllers.BindRoutes(ctx, router, service, repo, conf)

	go func() {
		if err := router.Run(); err != nil {
			log.Fatal("shutting down server...:", err.Error())
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := router; err != nil {
		log.Fatal(err)
	}
}

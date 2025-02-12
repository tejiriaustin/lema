package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	constants "github.com/tejiriaustin/lema/constants"
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
) error {
	router := gin.New()

	rateLimiter := middleware.NewRateLimiter(5, 10, time.Hour)

	router.Use(
		rateLimiter.RateLimit(),
		middleware.CORSMiddleware(),
		middleware.DefaultStructuredLogs(),
		middleware.ReadPaginationOptions(),
	)

	controllers.BindRoutes(ctx, router, service, repo, conf)

	srv := &http.Server{
		Addr:    conf.GetAsString(constants.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v\n", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	log.Printf("received shutdown signal: %v\n", sig)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return errors.Join(err, errors.New("server forced to shutdown"))
	}

	log.Println("server exited gracefully")
	return nil
}

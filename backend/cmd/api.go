// Package cmd /*
package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	constants "github.com/tejiriaustin/lema/constants"
	"github.com/tejiriaustin/lema/database"
	"github.com/tejiriaustin/lema/env"
	"github.com/tejiriaustin/lema/logger"
	"github.com/tejiriaustin/lema/models"
	"github.com/tejiriaustin/lema/repository"
	"github.com/tejiriaustin/lema/server"
	"github.com/tejiriaustin/lema/service"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long:  `API is a CLI backend application written in Go that empowers applications.`,
	Run:   startApi,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func startApi(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	lemaLogger, err := logger.NewProductionLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	config := setApiEnvironment()

	dbCfg := &database.Config{
		DB: config.GetAsString(constants.DB),
	}
	dbConn, err := database.Initialize(dbCfg)
	if err != nil {
		return
	}

	if config.GetAsString(constants.ShouldAutoMigrate) == "true" {
		tables := []interface{}{
			models.User{},
			models.Post{},
			models.Address{},
		}
		err = dbConn.Migrate(tables...)
		if err != nil {
			return
		}
	}

	rc := repository.NewRepositoryContainer(lemaLogger, dbConn)

	sc := service.NewService(lemaLogger, &config)

	server.Start(ctx, sc, rc, &config)
}

func setApiEnvironment() env.Environment {
	staticEnvironment := env.NewEnvironment()

	staticEnvironment.
		SetEnv(constants.Port, env.GetEnv(constants.Port, "8080")).
		SetEnv(constants.RedisDsn, env.MustGetEnv(constants.RedisDsn)).
		SetEnv(constants.DB, env.MustGetEnv(constants.DB)).
		SetEnv(constants.ShouldAutoMigrate, env.MustGetEnv(constants.ShouldAutoMigrate)).
		SetEnv(constants.JwtSecret, env.MustGetEnv(constants.JwtSecret)).
		SetEnv(constants.FrontendUrl, env.MustGetEnv(constants.FrontendUrl))

	return staticEnvironment
}

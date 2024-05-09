package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/vekaputra/tiger-kittens/cmd"
	"github.com/vekaputra/tiger-kittens/internal/app"
	"github.com/vekaputra/tiger-kittens/internal/config"
)

func main() {
	appConfig := config.Load()
	srv := app.NewServer(appConfig)

	// setup zerolog
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if appConfig.IsEnableDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	cmdConfig := cmd.Config{
		Use:   "tiger-kittens",
		Short: "API for tracking tiger populations",
		Long:  "API for tracking tiger populations based on list of tigers",
	}
	cmdDB := cmd.DB{
		Driver:   "postgres",
		Username: appConfig.DatabaseConfig.Username,
		Password: appConfig.DatabaseConfig.Password,
		Host:     appConfig.DatabaseConfig.Host,
		Port:     appConfig.DatabaseConfig.Port,
		Name:     appConfig.DatabaseConfig.DBName,
	}

	cli := cmd.New(cmdConfig, cmdDB, srv)
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}

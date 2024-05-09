package main

import (
	"github.com/vekaputra/tiger-kittens/cmd"
	"github.com/vekaputra/tiger-kittens/internal/config"
	"os"
)

func main() {
	appConfig := config.Load()

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
	cli := cmd.New(cmdConfig, cmdDB)
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}

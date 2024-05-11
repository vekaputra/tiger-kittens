package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vekaputra/tiger-kittens/pkg/database/sql"
)

func newDBSeedCommand(db DB) *cobra.Command {
	return &cobra.Command{
		Use:   "db:seed",
		Short: "seed the wrapper database",
		Run: func(cmd *cobra.Command, args []string) {
			absPath, err := filepath.Abs("db/seeds")
			if err != nil {
				log.Fatal(err)
			}
			dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.Username, db.Password, db.Name)
			dbSeed(absPath, dsn)
		},
	}
}

func dbSeed(dir, dsn string) {
	s, err := sql.NewSeed(dir, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Run(); err != nil {
		log.Fatal(err)
	}

	if err = s.Close(); err != nil {
		log.Fatal(err)
	}
}

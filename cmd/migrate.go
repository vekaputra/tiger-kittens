package cmd

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

type DB struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

func newDBMigrateCommand(db DB) *cobra.Command {
	return &cobra.Command{
		Use:   "db:migrate",
		Short: "Execute the migration from 'db/migrations' folder",
		Run: func(cmd *cobra.Command, args []string) {
			absPath, err := filepath.Abs("db/migrations")
			if err != nil {
				log.Fatal(err)
			}
			sourceURL := fmt.Sprintf("file:///%s", absPath)
			dbMigrate(sourceURL, db)
		},
	}
}

func dbMigrate(sourceURL string, db DB) {
	m, err := migrate.New(sourceURL, generateDatabaseURL(db))
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	srcErr, dbErr := m.Close()
	if srcErr != nil {
		log.Fatalln(srcErr)
	}
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
}

func generateDatabaseURL(db DB) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", db.Driver, db.Username, db.Password, db.Host, db.Port, db.Name)
}

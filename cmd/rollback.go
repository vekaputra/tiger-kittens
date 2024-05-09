package cmd

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"
)

func newDBRollbackCommand(db DB) *cobra.Command {
	return &cobra.Command{
		Use:   "db:rollback",
		Short: "rollback the migration to the previous 1 version",
		Run: func(cmd *cobra.Command, args []string) {
			absPath, err := filepath.Abs("db/migrations")
			if err != nil {
				log.Fatal(err)
			}
			sourceURL := fmt.Sprintf("file:///%s", absPath)
			step := -1
			dbRollback(sourceURL, db, step)
		},
	}
}

func dbRollback(sourceURL string, db DB, step int) {
	m, err := migrate.New(sourceURL, generateDatabaseURL(db))
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Steps(step); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	src, err := m.Close()
	if err != nil {
		log.Fatalln(err)
	}

	if src != nil {
		log.Fatalln(src)
	}
}

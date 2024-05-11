package sql

import (
	"fmt"
	"os"
	"sort"

	"github.com/jmoiron/sqlx"
)

type Seed struct {
	db    *sqlx.DB
	files []string
}

func NewSeed(dir, dsn string) (*Seed, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	filesInfo, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, fileInfo := range filesInfo {
		files = append(files, fmt.Sprintf("%s/%s", dir, fileInfo.Name()))
	}

	sort.Strings(files)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Seed{
		db:    db,
		files: files,
	}, nil
}

func (s *Seed) Close() error {
	return s.db.Close()
}

func (s *Seed) Run() error {
	for _, file := range s.files {
		b, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		if len(b) == 0 {
			continue
		}

		if _, err := s.db.Exec(string(b)); err != nil {
			return fmt.Errorf("unable to seed: '%s'. error: %w", file, err)
		}
	}

	return nil
}

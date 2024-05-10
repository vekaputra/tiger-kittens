package pgsql

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vekaputra/tiger-kittens/internal/config"
	pkgsqlx "github.com/vekaputra/tiger-kittens/pkg/database/sqlx"
)

var dbTest pkgsqlx.DBer

func generateDatabaseURL(dbConfig config.DatabaseConfig) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", "postgres", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
}

func migrateDB(dbConfig config.DatabaseConfig) {
	sourceURL := "file://../../../db/migrations"
	m, err := migrate.New(sourceURL, generateDatabaseURL(dbConfig))
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	srcErr, dbErr := m.Close()
	if srcErr != nil {
		panic(srcErr)
	}
	if dbErr != nil {
		panic(dbErr)
	}
}

func createDB() (pkgsqlx.DBer, func()) {
	dbName := fmt.Sprintf("test_%s", uuid.New().String())
	dbName = strings.ReplaceAll(dbName, "-", "_")

	dbConfig := config.DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "root",
		DBName:   dbName,
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)); err != nil {
		panic(err)
	}

	dsnTest := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DBName)
	dbTest, err = sqlx.Open("postgres", dsnTest)
	if err != nil {
		panic(err)
	}

	if _, err = dbTest.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`); err != nil {
		panic(err)
	}
	migrateDB(dbConfig)

	return dbTest, func() {
		dbTest.Close()
		dropDB(dsn, dbName)
	}
}

func dropDB(dsn, dbName string) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err = db.Exec(fmt.Sprintf(`DROP DATABASE %s;`, dbName)); err != nil {
		panic(err)
	}
}

func TestMain(t *testing.M) {
	createdDB, closeDB := createDB()
	dbTest = createdDB

	code := t.Run()
	closeDB()
	os.Exit(code)
}

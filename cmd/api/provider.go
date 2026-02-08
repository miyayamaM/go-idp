package main

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun/driver/pgdriver"
)

func ProvideDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/go_idp?sslmode=disable"
	}
	return sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn))), nil
}

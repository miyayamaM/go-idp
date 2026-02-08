package main

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func ProvideDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/go_idp?sslmode=disable"
	}
	return sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn))), nil
}

func ProvideBunDB(db *sql.DB) *bun.DB {
	return bun.NewDB(db, pgdialect.New())
}

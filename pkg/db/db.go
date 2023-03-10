package db

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

func NewDB() (*bun.DB, error) {
	pgSQLdb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(os.Getenv("DB_CONN_URL")),
	))
	err := pgSQLdb.Ping()
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(pgSQLdb, pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

package postgres

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func NewPgConn() (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://postgres:%v@localhost:5432/postgres?sslmode=disable", os.Getenv("POSTGRES_PASSWORD"))

	conn, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

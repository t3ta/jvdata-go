package jvdata

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

type DBManager struct {
	Conn *pgx.Conn
}

func (db *DBManager) connect() (err error) {
	db.Conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	return
}

func (db *DBManager) close() {
	db.Conn.Close(context.Background())
}

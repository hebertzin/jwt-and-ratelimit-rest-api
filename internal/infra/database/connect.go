package database

import (
	"context"
	"database/sql"
	"log"
)

type Connection struct {
	DNS string
}

func (config *Connection) MustConnect(ctx context.Context) *sql.DB {
	db, err := sql.Open("pgx", config.DNS)
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("cannot connect to postgres:", err)
	}

	return db
}

package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	routing "github.com/jwt-and-ratelimit-rest-api/src/router"
)

func main() {

	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("error connecting database", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("cannot connect to postgres", err)
	}

	r := chi.NewRouter()
	runMigrations(db)
	routing.UsersGoupRouter(r, db)

	http.ListenAndServe(":8080", r)
}

func runMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("error creating migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("error creating migration instance:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error running migrations:", err)
	}

	log.Println("migrations applied successfully")
}

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
	_ "github.com/jwt-and-ratelimit-rest-api/docs"
	"github.com/jwt-and-ratelimit-rest-api/src/middlewares"
	routing "github.com/jwt-and-ratelimit-rest-api/src/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           JWT and Rate Limit REST API
// @version         1.0
// @description     A RESTful API built with Go, implementing authentication with JWT and request rate limiting following industry best practices.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://www.hebertzin.com
// @contact.email  hebertsantosdeveloper@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token in the format: Bearer <token>
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

	r.Use(middlewares.RateLimitMiddleware)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	runMigrations(db)

	routing.UsersGoupRouter(r, db)

	routing.AuthenticationGroupRouter(r, db)

	if err = http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error starting http server", err)
	}
}

func runMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("error creating migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
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

package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jwt-and-ratelimit-rest-api/docs"
	"github.com/jwt-and-ratelimit-rest-api/packages/middlewares"
	routing "github.com/jwt-and-ratelimit-rest-api/packages/router"
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
		log.Fatal("error opening database:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("cannot connect to postgres:", err)
	}
	cancel()

	r := chi.NewRouter()

	r.Use(middlewares.RateLimitMiddleware)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	runMigrations(db)

	routing.UsersGroupRouter(r, db)

	routing.AuthenticationGroupRouter(r, db)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("server forced to shutdown:", err)
	}

	if err := db.Close(); err != nil {
		log.Println("error closing database:", err)
	}

	log.Println("server exited")
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

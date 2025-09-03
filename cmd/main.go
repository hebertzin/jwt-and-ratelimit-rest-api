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
	_ "github.com/hebertzin/jwt-and-ratelimit-rest-api/docs"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/infra/database"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/middlewares"
	routing "github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/router"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"

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
	buildEnvs()

	r := chi.NewRouter()

	r.Use(middlewares.RateLimitMiddleware)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	conn := database.Connection{DNS: os.Getenv("DATABASE_URL")}

	db := conn.MustConnect(context.Background())

	runMigrations(db)

	buildRoutes(r, db)

	srv := buildServer(r)

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

func buildEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, relying on system environment variables")
	}
}

func buildRoutes(r chi.Router, db *sql.DB) {
	routing.UsersGroupRouter(r, db)
	routing.AuthenticationGroupRouter(r, db)
}

func buildServer(r http.Handler) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	return srv
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

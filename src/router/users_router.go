package routing

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/jwt-and-ratelimit-rest-api/src/handler"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/security"
	"github.com/jwt-and-ratelimit-rest-api/src/middlewares"
	"github.com/jwt-and-ratelimit-rest-api/src/services"
	"github.com/jwt-and-ratelimit-rest-api/src/utils/validation"
)

func UsersGoupRouter(r chi.Router, db *sql.DB) {
	userRepository := repository.NewUsersRepository(db)

	hasher := security.NewBcryptHasher(10)

	pv := validation.NewPayloadValidate()

	userService := services.NewUserService(userRepository, hasher, pv)

	userHandler := handler.NewUserHandler(userService)

	r.Use(middlewares.DoFilter)

	r.Post("/api/v1/accounts", userHandler.Create)
}

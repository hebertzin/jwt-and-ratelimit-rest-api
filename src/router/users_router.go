package routing

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/jwt-and-ratelimit-rest-api/src/handler"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/security"
	"github.com/jwt-and-ratelimit-rest-api/src/services"
)

func UsersGoupRouter(r chi.Router, db *sql.DB) {
	userRepository := repository.NewUsersRepository(db)
	hasher := security.NewBcryptHasher(10)
	userService := services.NewUserService(userRepository, hasher)
	userHandler := handler.NewUserHanlder(userService)

	r.Post("/api/v1/users", userHandler.Create)

}

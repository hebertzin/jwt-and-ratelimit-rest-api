package factory

import (
	"database/sql"

	"github.com/jwt-and-ratelimit-rest-api/src/handler"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/security"
	"github.com/jwt-and-ratelimit-rest-api/src/services"
)

func AuthenticationFactory(db *sql.DB) *handler.AuthenticationHandler {
	r := repository.NewUsersRepository(db)

	h := security.NewBcryptHasher(10)

	s := services.NewAuthenticationService(r, h)

	return handler.NewAuthenticatorHandler(s)
}

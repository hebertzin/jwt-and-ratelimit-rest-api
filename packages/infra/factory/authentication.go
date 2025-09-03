package factory

import (
	"database/sql"

	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/handler"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/infra/repository"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/infra/security"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/services"
)

func AuthenticationFactory(db *sql.DB) *handler.AuthenticationHandler {
	r := repository.NewUsersRepository(db)

	h := security.NewBcryptHasher(10)

	s := services.NewAuthenticationService(r, h)

	return handler.NewAuthenticatorHandler(s)
}

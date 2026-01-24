package factory

import (
	"database/sql"

	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/handler"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/infra/repository"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/infra/security"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/services"
)

func AuthenticationFactory(db *sql.DB) *handler.AuthenticationHandler {
	r := repository.NewUsersRepository(db)

	h := security.NewBcryptHasher(10)

	s := services.NewAuthenticationService(r, h)

	return handler.NewAuthenticatorHandler(s)
}

package factory

import (
	"database/sql"

	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/handler"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/infra/repository"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/infra/security"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/services"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/internal/utils/validation"
)

func UsersFactory(db *sql.DB) *handler.UserHandler {
	r := repository.NewUsersRepository(db)

	h := security.NewBcryptHasher(10)

	v := validation.NewPayloadValidate()

	s := services.NewUserService(r, h, v)

	return handler.NewUserHandler(s)

}

package factory

import (
	"database/sql"

	"github.com/jwt-and-ratelimit-rest-api/src/handler"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/security"
	"github.com/jwt-and-ratelimit-rest-api/src/services"
	"github.com/jwt-and-ratelimit-rest-api/src/utils/validation"
)

func UsersFactory(db *sql.DB) *handler.UserHandler {
	r := repository.NewUsersRepository(db)

	h := security.NewBcryptHasher(10)

	v := validation.NewPayloadValidate()

	s := services.NewUserService(r, h, v)

	return handler.NewUserHandler(s)

}

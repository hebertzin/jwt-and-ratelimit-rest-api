package factory

import (
	"database/sql"

	"github.com/jwt-and-ratelimit-rest-api/packages/handler"
	"github.com/jwt-and-ratelimit-rest-api/packages/infra/repository"
	"github.com/jwt-and-ratelimit-rest-api/packages/infra/security"
	"github.com/jwt-and-ratelimit-rest-api/packages/services"
	"github.com/jwt-and-ratelimit-rest-api/packages/utils/validation"
)

func UsersFactory(db *sql.DB) *handler.UserHandler {
	r := repository.NewUsersRepository(db)

	h := security.NewBcryptHasher(10)

	v := validation.NewPayloadValidate()

	s := services.NewUserService(r, h, v)

	return handler.NewUserHandler(s)

}

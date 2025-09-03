package routing

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/infra/factory"
)

func UsersGroupRouter(r chi.Router, db *sql.DB) {
	u := factory.UsersFactory(db)

	r.Post("/api/v1/accounts", u.Create)
}

package routing

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/factory"
)

func UsersGoupRouter(r chi.Router, db *sql.DB) {
	u := factory.UsersFactory(db)

	r.Post("/api/v1/accounts", u.Create)
}

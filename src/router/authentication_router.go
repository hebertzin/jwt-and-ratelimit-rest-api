package routing

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/jwt-and-ratelimit-rest-api/src/infra/factory"
)

func AuthenticationGroupRouter(r chi.Router, db *sql.DB) {
	a := factory.AuthenticationFactory(db)

	r.Post("/api/v1/authentication", a.Authenticate)
}

package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hebertzin/jwt-and-ratelimit-rest-api/packages/infra/security"
)

func DoFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"error":"missing authorization header"}`)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"error":"invalid authorization header"}`)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if err := security.VerifyToken(tokenString); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"error":"invalid token"}`)
			return
		}

		next.ServeHTTP(w, r)
	})
}

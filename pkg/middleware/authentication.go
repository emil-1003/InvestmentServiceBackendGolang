package middleware

import (
	"net/http"

	"github.com/emil-1003/InvestmentServiceBackendGolang/pkg/authentication"
)

func Secure(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Check if token is valid
		_, ok := authentication.GetToken(r)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

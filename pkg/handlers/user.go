package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emil-1003/InvestmentServiceBackendGolang/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var body struct {
			Name     string
			Email    string
			Password string
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, fmt.Errorf("failed to read body: %w", err).Error(), http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
		if err != nil {
			http.Error(w, fmt.Errorf("failed to hash password: %w", err).Error(), http.StatusBadRequest)
			return
		}

		if err = models.CreateUser(body.Name, body.Email, hashedPassword); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte("user was registered successfully"))
	}
}

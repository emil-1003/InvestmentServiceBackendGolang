package models

import (
	"fmt"

	"github.com/emil-1003/InvestmentServiceBackendGolang/pkg/database"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
	Created  string `json:"created"`
	Login    string `json:"login"`
}

func CreateUser(name string, email string, hashedPassword []byte) error {
	_, err := database.DB.Exec(`
		INSERT INTO users (name, email, password, role_id)
		VALUES (?, ?, ?, 1)
	`, name, email, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return err
}

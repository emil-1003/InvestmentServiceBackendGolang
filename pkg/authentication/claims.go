package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/emil-1003/InvestmentServiceBackendGolang/pkg/models"
)

type Claims struct {
	Uid  int    `json:"uid"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// 1 hour
var maxAge = time.Minute * 60

func CreateToken(user models.User) (string, error) {
	claims := &Claims{
		Uid:  user.ID,
		Role: user.Role.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(maxAge).Unix(), // Token will expire after 10 minutes
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		return "", fmt.Errorf("while creating jwt: %w", err)
	}

	return tokenString, nil
}

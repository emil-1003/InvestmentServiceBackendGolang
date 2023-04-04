package authentication

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET_WORD")))
	if err != nil {
		return "", fmt.Errorf("while creating jwt: %w", err)
	}

	return tokenString, nil
}

// Returns parsed token and bool thats true if token is valid
func GetToken(r *http.Request) (*jwt.Token, bool) {
	// Get the value of the Authorization header from the HTTP request
	authHeader := r.Header.Get("Authorization")

	// Check that the Authorization header is not empty and starts with the string "Bearer "
	if authHeader == "" && !strings.HasPrefix(authHeader, "Bearer") {
		return nil, false
	}

	// If the Authorization header is valid, extract the token from it by removing the "Bearer " prefix
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET_WORD")), nil
	})
	if err != nil {
		fmt.Println(fmt.Errorf("while parsing token: %v", err).Error())
		return nil, false
	}

	return token, token.Valid
}

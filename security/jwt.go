package security

//Reference link: https://thedevelopercafe.com/articles/jwt-with-go-52d6bbcaa2bf

import (
	"example/backend-github-trending/model"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "SECRET KEY"

func GenToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtCustomClaims{
		UserId: user.UserId,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{},
	})

	// Create the actual JWT token

	signedString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}

	return signedString, nil
}
package security

//Reference link: https://thedevelopercafe.com/articles/jwt-with-go-52d6bbcaa2bf

import (
	"example/backend-github-trending/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "SECRET KEY"

func GenToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtCustomClaims{
		UserId: user.UserId,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{},
		Date: time.Now(),
	})

	// Create the actual JWT token

	signedString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", fmt.Errorf("error creating signed string: %v", err)
	}

	return signedString, nil
}

func ParseToken(jwtToken string) (model.JwtCustomClaims, error) {
	var customClaim model.JwtCustomClaims

	token, err := jwt.ParseWithClaims(jwtToken, &customClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return model.JwtCustomClaims{}, err
	}

	if !token.Valid {
		return model.JwtCustomClaims{}, err
	}

	return customClaim, nil
}
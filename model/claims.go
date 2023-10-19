package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JwtCustomClaims struct {
	UserId string
	Role string
	jwt.RegisteredClaims
	Date time.Time
}
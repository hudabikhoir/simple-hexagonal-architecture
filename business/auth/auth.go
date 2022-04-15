package auth

import "github.com/golang-jwt/jwt"

type Auth struct {
	Token string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

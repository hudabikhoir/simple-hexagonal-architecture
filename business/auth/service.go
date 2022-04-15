package auth

import (
	"clean-hexa/business/auth/spec"
	"clean-hexa/config"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type Service interface {
	Login(input spec.InputLogin) (auth *Auth, err error)
}

type service struct {
	validate *validator.Validate
	config   *config.AppConfig
}

// NewService ...
func NewService(config *config.AppConfig) Service {
	return &service{
		validate: validator.New(),
		config:   config,
	}
}

var users = map[string]string{
	"admin1": "passadmin1",
	"admin2": "passadmin2",
}

func (s *service) Login(input spec.InputLogin) (auth *Auth, err error) {
	excepctedPassword, isExist := users[input.Username]

	if !isExist || excepctedPassword != input.Password {
		return nil, errors.New("Invalid username or password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: input.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	screetKey := []byte(s.config.App.JWTKey)

	tokenString, err := token.SignedString(screetKey)
	if err != nil {
		return nil, err
	}
	auth = &Auth{
		Token: tokenString,
	}

	return auth, nil
}

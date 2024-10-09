package jwt

import (
	"errors"
	"fmt"
	"time"

	jwtService "github.com/golang-jwt/jwt/v5"
	errs "github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
)

type JWT interface {
	Create() (string, error)
	Parse(tokenString string) error
}

type jwt struct {
	HMACSecret []byte
	Expiry     time.Duration
}

func New(hmacSecret []byte, expiry time.Duration) JWT {
	return &jwt{HMACSecret: hmacSecret, Expiry: expiry}
}

func (j *jwt) Create() (string, error) {
	token := jwtService.NewWithClaims(jwtService.SigningMethodHS256, jwtService.RegisteredClaims{
		ExpiresAt: jwtService.NewNumericDate(time.Now().Add(j.Expiry)),
	})

	tokenString, err := token.SignedString(j.HMACSecret)

	return tokenString, err
}

func (j *jwt) Parse(tokenString string) error {
	token, err := jwtService.Parse(tokenString, func(token *jwtService.Token) (any, error) {
		if _, ok := token.Method.(*jwtService.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.HMACSecret, nil
	})

	switch {
	case token.Valid:
		return nil
	case errors.Is(err, jwtService.ErrTokenExpired) || errors.Is(err, jwtService.ErrTokenNotValidYet):
		return errs.ErrorExpiredToken
	default:
		return errs.ErrorInvalidToken
	}
}

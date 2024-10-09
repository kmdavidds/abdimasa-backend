package service

import (
	"os"

	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/jwt"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

type AuthService interface {
	Login(req dto.LoginRequest) (string, error)
}

type authService struct {
	val validator.Validator
	jwt jwt.JWT
}

func NewAuthService(
	val validator.Validator,
	jwt jwt.JWT,
) AuthService {
	return &authService{val, jwt}
}

func (as *authService) Login(req dto.LoginRequest) (string, error) {
	valErr := as.val.Validate(req)
	if valErr != nil {
		return "", valErr
	}

	if req.ID != os.Getenv("ADMIN_ID") || req.Password != os.Getenv("ADMIN_PASSWORD") {
		return "", errors.ErrorInvalidCredentials
	}

	tokenString, err := as.jwt.Create()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
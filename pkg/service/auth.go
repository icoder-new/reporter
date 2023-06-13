package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
	"github.com/icoder-new/reporter/utils"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	pwd, err := generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = pwd
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	password, err := generatePasswordHash(password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(utils.AppSettings.AppParams.TokenTTL) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
	})

	return token.SignedString([]byte(utils.AppSettings.AppParams.SecretKey))
}

func (s *AuthService) ParseToken(token string) (int, error) {
	_token, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, utils.ErrInvalidSigningKey
		}
		return []byte(utils.AppSettings.AppParams.SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := _token.Claims.(*tokenClaims)
	if !ok {
		return 0, utils.ErrInvalidTypeOfClaims
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

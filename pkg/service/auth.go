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

func (s *AuthService) CreateUser(firstname, lastname, username, email, password string) (int, error) {
	var user models.User
	user.Firstname = firstname
	user.Lastname = lastname
	user.Username = username
	user.Email = email
	user.IsActive = true

	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	user.Password = string(pwd)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, username, password string) (string, error) {
	user, err := s.repo.GetUser(email, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(utils.AppSettings.AppParams.TokenTTL) * time.Minute),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		user.ID,
	})
	t, err := token.SignedString([]byte(utils.AppSettings.AppParams.SecretKey))
	return t, err
}

func (s *AuthService) ParseToken(token string) (int, error) {
	_token, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return -1, utils.ErrInvalidSigningKey
		}
		return []byte(utils.AppSettings.AppParams.SecretKey), nil
	})
	if err != nil {
		return -1, err
	}

	claims, ok := _token.Claims.(*tokenClaims)
	if !ok {
		return -1, utils.ErrInvalidTypeOfClaims
	}

	return claims.UserID, nil
}

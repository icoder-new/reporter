package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func GenerateToken(username, password string) (string, error) {
	/* password, err := GeneratePasswordHash(password)
	if err != nil {
	return "", err
	} */

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(AppSettings.AppParams.TokenTTL) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		4, // TODO: change it into user.ID
	})

	return token.SignedString([]byte(AppSettings.AppParams.SecretKey))
}

func ParseToken(token string) (int, error) {
	_token, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningKey
		}
		return []byte(AppSettings.AppParams.SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := _token.Claims.(*tokenClaims)
	if !ok {
		return 0, ErrInvalidTypeOfClaims
	}

	return claims.UserId, nil
}

func GeneratePasswordHash(password string) (string, error) {
	hashedPasswod, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return cast.ToString(hashedPasswod), nil
}

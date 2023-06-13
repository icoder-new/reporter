package utils

import "errors"

var (
	ErrInvalidSigningKey   = errors.New("invalid signing method")
	ErrInvalidTypeOfClaims = errors.New("token claims are not type of `*tokenClaims`")
	ErrUserIdNotFound      = errors.New("user id not found")
)

package service

import (
	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserById(id int) (models.User, error)
}

type Service struct {
	Authorization
	User
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		User:          NewUserService(repository.User),
	}
}

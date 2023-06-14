package service

import (
	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
)

type Authorization interface {
	CreateUser(fistname, lastname, username, email, password string) (int, error)
	GenerateToken(email, username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserById(id int) (models.User, error)
	UpdateUser(id int, firstname, lastname, email, username, password string) (models.User, error)
	DeleteUserById(id int) error
	RestoreUserById(id int) error
}

type Account interface {
	CreateAccountByUserId(account models.Account) (int, error)
	UpdateAccountByUserId(account models.Account) (models.Account, error)
	DeleteAccountByUserId(userId int) error
	RestoreAccountByUserId(userId int) error
}

type Service struct {
	Authorization
	User
	Account
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		User:          NewUserService(repository.User),
		Account:       NewAccountService(repository.Account),
	}
}

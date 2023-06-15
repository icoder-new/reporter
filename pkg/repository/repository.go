package repository

import (
	"github.com/icoder-new/reporter/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, username string) (models.User, error)
}

type User interface {
	GetUserById(id int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUserById(id int) error
	RestoreUserById(id int) error
}

type Account interface {
	GetAccountsByUserId(userId int) ([]models.Account, error)
	GetAccountUserId(id, userId int) (models.Account, error)
	CreateAccountByUserId(account models.Account) (int, error)
	UpdateAccountByUserId(account models.Account) (models.Account, error)
	DeleteAccountByUserId(id, userId int) error
	RestoreAccountByUserId(id, userId int) error
}

type Repository struct {
	Authorization
	User
	Account
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserRepository(db),
		Account:       NewAccountRepository(db),
	}
}

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
}

type Repository struct {
	Authorization
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserRepository(db),
	}
}

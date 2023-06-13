package repository

import (
	"github.com/icoder-new/reporter/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (p *AuthPostgres) CreateUser(user models.User) (int, error) {
	if err := p.db.Create(&user).Error; err != nil {
		return 0, nil
	}

	return user.ID, nil
}

func (p *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var u models.User
	err := p.db.Where("username = ? AND password = ?", username, password).First(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

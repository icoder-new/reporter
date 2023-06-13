package repository

import (
	"github.com/icoder-new/reporter/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := u.db.Model(models.User{}).Where("id = ? AND is_active = TRUE", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

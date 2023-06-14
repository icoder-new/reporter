package repository

import (
	"github.com/icoder-new/reporter/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) CreateAccountByUserId(account models.Account) (int, error) {
	if err := r.db.Create(&account).Error; err != nil {
		return -1, err
	}

	return account.ID, nil
}

func (r *AccountRepository) UpdateAccountByUserId(account models.Account) (models.Account, error) {
	if err := r.db.Save(&account).Error; err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func (r *AccountRepository) DeleteAccountByUserId(userId int) error {
	return r.db.Model(&models.Account{}).Where("user_id = ? AND is_active = ?", userId, true).Update("is_active", false).Error
}

func (r *AccountRepository) RestoreAccountByUserId(userId int) error {
	return r.db.Model(&models.Account{}).Where("user_id = ? AND is_active = ?", userId, false).Update("is_active", true).Error
}

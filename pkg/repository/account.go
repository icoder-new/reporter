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

func (r *AccountRepository) GetAccountsByUserId(userId int) ([]models.Account, error) {
	var accounts []models.Account

	err := r.db.Where("user_id = ? AND is_active = ?", userId, true).Find(&accounts).Error
	if err != nil {
		return accounts, err
	}

	return accounts, nil
}

func (r *AccountRepository) GetAccountUserId(id, userId int) (models.Account, error) {
	var account models.Account
	err := r.db.Where("id = ? AND user_id = ? AND is_active = ?", id, userId, true).First(&account).Error
	if err != nil {
		return account, err
	}

	return account, nil
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

func (r *AccountRepository) DeleteAccountByUserId(id, userId int) error {
	return r.db.Model(&models.Account{}).Where(
		"id = ? AND user_id = ? AND is_active = ?", id, userId, true).Update(
		"is_active", false,
	).Error
}

func (r *AccountRepository) RestoreAccountByUserId(id, userId int) error {
	return r.db.Model(&models.Account{}).Where(
		"id = ? AND user_id = ? AND is_active = ?", id, userId, false).Update(
		"is_active", true,
	).Error
}

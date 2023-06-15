package service

import (
	"time"

	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
	"github.com/icoder-new/reporter/utils"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) GetAccount(id, userId int) (models.Account, error) {
	return s.repo.GetAccountUserId(id, userId)
}

func (s *AccountService) GetAllAccounts(userId int) ([]models.Account, error) {
	return s.repo.GetAccountsByUserId(userId)
}

func (s *AccountService) CreateAccount(userId int, name string, balance float64) (int, error) {
	var account models.Account

	account.UserID = userId

	if utils.CheckField(name) {
		account.Name = name
	}

	if err := utils.CheckBalance(balance); err == nil {
		account.Balance = balance
	} else {
		return -1, err
	}

	account.IsActive = true

	return s.repo.CreateAccountByUserId(account)
}

func (s *AccountService) UpdateAccount(id, userId int, name string, balance float64) (models.Account, error) {
	var account models.Account
	account, err := s.repo.GetAccountUserId(id, userId)
	if err != nil {
		return account, err
	}

	if utils.CheckField(name) {
		account.Name = name
	} else {
		return account, utils.ErrInvalidAccountName
	}

	if err := utils.CheckBalance(balance); err == nil {
		account.Balance = balance
	} else {
		return account, err
	}

	account.UpdatedAt = time.Now()

	return s.repo.UpdateAccountByUserId(account)
}

func (s *AccountService) DeleteAccount(id, userId int) error {
	return s.repo.DeleteAccountByUserId(id, userId)
}

func (s *AccountService) RestoreAccount(id, userId int) error {
	return s.repo.RestoreAccountByUserId(id, userId)
}

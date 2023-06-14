package service

import (
	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) CreateAccountByUserId(account models.Account) (int, error) {
	return -1, nil
}

func (s *AccountService) UpdateAccountByUserId(account models.Account) (models.Account, error) {
	return models.Account{}, nil
}

func (s *AccountService) DeleteAccountByUserId(userId int) error {
	return nil
}

func (s *AccountService) RestoreAccountByUserId(userId int) error {
	return nil
}

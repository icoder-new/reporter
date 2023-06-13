package service

import (
	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUserById(id int) (models.User, error) {
	return u.repo.GetUserById(id)
}

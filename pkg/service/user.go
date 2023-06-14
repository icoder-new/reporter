package service

import (
	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserService) DeleteUserById(id int) error {
	return u.repo.DeleteUserById(id)
}

func (u *UserService) RestoreUserById(id int) error {
	return u.repo.RestoreUserById(id)
}

func (u *UserService) UpdateUser(id int, firstname, lastname, email, username, password string) (models.User, error) {
	user, err := u.GetUserById(id)
	if err != nil {
		return user, err
	}

	if checkField(firstname) {
		user.Firstname = firstname
	}

	if checkField(lastname) {
		user.Lastname = lastname
	}

	if checkField(email) {
		user.Email = email
	}

	if checkField(username) {
		user.Username = username
	}

	if isChangeablePassword(user.Password, password) == nil {
		pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return user, err
		}

		user.Password = string(pwd)
	}

	return u.repo.UpdateUser(user)
}

func checkField(field string) bool {
	if field == "" || field == " " || len(field) > 50 {
		return false
	}

	return true
}

func isChangeablePassword(userPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
}

package service

import (
	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
)

type Authorization interface {
	CreateUser(fistname, lastname, username, email, password string) (int, error)
	GenerateToken(email, username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserById(id int) (models.User, error)
	UpdateUser(id int, firstname, lastname, email, username, password string) (models.User, error)
	DeleteUserById(id int) error
	RestoreUserById(id int) error
	UpdatePictureUser(id int, filepath string) (models.User, error)
	UploadUserPicture(id int, filepath string) (models.User, error)
}

type Account interface {
	GetAccount(id, userId int) (models.Account, error)
	GetAllAccounts(userId int) ([]models.Account, error)
	CreateAccount(userId int, name string, balance float64) (int, error)
	UpdateAccount(id, userId int, name string, balance float64) (models.Account, error)
	DeleteAccount(id, userId int) error
	RestoreAccount(id, userId int) error
	ChangePictureAccount(id, userId int, filepath string) (models.Account, error)
	UploadAccountPicture(id, userId int, filepath string) (models.Account, error)
}

type Category interface {
	CreateCategory(name, description string) (models.Category, error)
	GetCategories() ([]models.Category, error)
	GetCategory(id int) (models.Category, error)
	UpdateCategory(id int, name, description string) (models.Category, error)
	UploadPictureCategory(id int, filepath string) (models.Category, error)
	ChangePictureCategory(id int, filepath string) (models.Category, error)
	DeleteCategory(id int) error
	RestoreCategory(id int) error
}

type Product interface {
	CreateProduct(catId int, name, description string, price float64) (models.Product, error)
	GetProducts(catId int) ([]models.Product, error)
	GetProduct(id, catId int) (models.Product, error)
	UpdateProduct(id, catId int, name, description string, price float64) (models.Product, error)
	UploadPictureProduct(id, catId int, filepath string) (models.Product, error)
	ChangePictureProduct(id, catId int, filepath string) (models.Product, error)
	DeleteProduct(id, catId int) error
	RestoreProduct(id, catId int) error
}

type Service struct {
	Authorization
	User
	Account
	Category
	Product
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		User:          NewUserService(repository.User),
		Account:       NewAccountService(repository.Account),
	}
}

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
	UpdateUser(user models.User) (models.User, error)
	DeleteUserById(id int) error
	RestoreUserById(id int) error
}

type Account interface {
	GetAccounts(userId int) ([]models.Account, error)
	GetAccount(id, userId int) (models.Account, error)
	CreateAccount(account models.Account) (int, error)
	UpdateAccount(account models.Account) (models.Account, error)
	DeleteAccount(id, userId int) error
	RestoreAccount(id, userId int) error
}

type Category interface {
	CreateCategory(cat models.Category) (models.Category, error)
	GetCategories() ([]models.Category, error)
	GetCategory(id int) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(id int) error
	RestoreCategory(id int) error
}

type Product interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetProducts(catId int) ([]models.Product, error)
	GetProduct(id, catId int) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id, catId int) error
	RestoreProduct(id, catId int) error
}

type Repository struct {
	Authorization
	User
	Account
	Category
	Product
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserRepository(db),
		Account:       NewAccountRepository(db),
		Category:      NewCategoryRepository(db),
		Product:       NewProductRepository(db),
	}
}

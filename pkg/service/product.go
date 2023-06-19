package service

import (
	"fmt"
	"os"
	"time"

	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
	"github.com/icoder-new/reporter/utils"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) CreateProduct(catId int, name, description string, price float64) (models.Product, error) {
	var product models.Product

	product.CategoryID = catId

	if utils.CheckField(name) {
		product.Name = name
	} else {
		return product, utils.ErrInvalidName
	}

	if utils.CheckField(description) {
		product.Description = description
	} else {
		return product, utils.ErrDescription
	}

	if err := utils.CheckBalance(price); err != nil {
		return product, err
	}

	product.Price = price

	return p.repo.CreateProduct(product)
}

func (p *ProductService) GetProducts(catId int) ([]models.Product, error) {
	return p.repo.GetProducts(catId)
}

func (p *ProductService) GetProduct(id, catId int) (models.Product, error) {
	return p.repo.GetProduct(id, catId)
}

func (p *ProductService) UpdateProduct(id, catId int, name, description string, price float64) (models.Product, error) {
	product, err := p.GetProduct(id, catId)
	if err != nil {
		return product, err
	}

	if utils.CheckField(name) {
		product.Name = name
	}

	if utils.CheckField(description) {
		product.Description = description
	}

	if utils.CheckBalance(price) == nil {
		product.Price = price
	}

	return p.repo.UpdateProduct(product)
}

func (p *ProductService) UploadPictureProduct(id, catId int, filepath string) (models.Product, error) {
	product, err := p.GetProduct(id, catId)
	if err != nil {
		return product, err
	}

	product.Picture = filepath
	product.UpdatedAt = time.Now()

	return p.repo.UpdateProduct(product)
}

func (p *ProductService) ChangePictureProduct(id, catId int, filepath string) (models.Product, error) {
	product, err := p.GetProduct(id, catId)
	if err != nil {
		return product, err
	}

	if err := os.Remove(fmt.Sprintf("./files/layouts/%s", product.Picture)); err != nil {
		return product, err
	}

	product.Picture = filepath
	product.UpdatedAt = time.Now()

	return p.repo.UpdateProduct(product)
}

// TODO
func (p *ProductService) DeleteProduct(id, catId int) error {
	return p.repo.DeleteProduct(id, catId)
}

// TODO
func (p *ProductService) RestoreProduct(id, catId int) error {
	return p.repo.RestoreProduct(id, catId)
}

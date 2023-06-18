package repository

import (
	"github.com/icoder-new/reporter/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) CreateProduct(product models.Product) (models.Product, error) {
	if err := p.db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) GetProducts(catId int) ([]models.Product, error) {
	var products []models.Product
	if err := p.db.
		Model(&models.Product{}).
		Where("category_id = ?", catId).
		Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p *ProductRepository) GetProduct(id, catId int) (models.Product, error) {
	var product models.Product
	if err := p.db.
		Model(&models.Product{}).
		Where("id = ? AND category_id = ?", id, catId).
		First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) UpdateProduct(product models.Product) (models.Product, error) {
	if err := p.db.Model(&models.Product{}).Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) DeleteProduct(id, catId int) error {
	return p.db.Model(&models.Product{}).Update("is_active = ?", false).Error
}

func (p *ProductRepository) RestoreProduct(id, catId int) error {
	return p.db.Model(&models.Product{}).Update("is_active = ?", true).Error
}

package service

import (
	"fmt"
	"os"
	"time"

	"github.com/icoder-new/reporter/models"
	"github.com/icoder-new/reporter/pkg/repository"
	"github.com/icoder-new/reporter/utils"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) CreateCategory(name, description string) (models.Category, error) {
	var cat models.Category

	if utils.CheckField(name) {
		cat.Name = name
	} else {
		return cat, utils.ErrInvalidName
	}

	if utils.CheckField(description) {
		cat.Description = description
	} else {
		return cat, utils.ErrDescription
	}

	return c.repo.CreateCategory(cat)
}

func (c *CategoryService) GetCategories() ([]models.Category, error) {
	return c.repo.GetCategories()
}

func (c *CategoryService) GetCategory(id int) (models.Category, error) {
	return c.repo.GetCategory(id)
}

func (c *CategoryService) UpdateCategory(id int, name, description string) (models.Category, error) {
	cat, err := c.GetCategory(id)
	if err != nil {
		return cat, err
	}

	if utils.CheckField(name) {
		cat.Name = name
	}

	if utils.CheckField(description) {
		cat.Description = description
	}

	cat.UpdatedAt = time.Now()

	return c.repo.UpdateCategory(cat)
}

func (c *CategoryService) UploadPictureCategory(id int, filepath string) (models.Category, error) {
	cat, err := c.GetCategory(id)
	if err != nil {
		return cat, err
	}

	cat.Picture = filepath
	cat.UpdatedAt = time.Now()

	return c.repo.UpdateCategory(cat)
}

func (c *CategoryService) ChangePictureCategory(id int, filepath string) (models.Category, error) {
	cat, err := c.GetCategory(id)
	if err != nil {
		return cat, err
	}

	if err := os.Remove(fmt.Sprintf("./files/layouts/%s", cat.Picture)); err != nil {
		return cat, err
	}

	cat.Picture = filepath
	cat.UpdatedAt = time.Now()

	return c.repo.UpdateCategory(cat)
}

// TODO
func (c *CategoryService) DeleteCategory(id int) error {
	return c.repo.DeleteCategory(id)
}

// TODO
func (c *CategoryService) RestoreCategory(id int) error {
	return c.repo.RestoreCategory(id)
}

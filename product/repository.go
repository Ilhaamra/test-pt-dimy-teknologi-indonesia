package product

import (
	"test-dimy-teknologi/models"

	"gorm.io/gorm"
)

type RepositoryProduct interface {
	FindAllProduct() ([]models.Product, error)
	FindProduct(id int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

type repositoryProductImpl struct {
	DB *gorm.DB
}

func NewRespositoryProduct(db *gorm.DB) RepositoryProduct {
	return &repositoryProductImpl{
		DB: db,
	}
}

// CreateProduct implements RepositoryProduct.
func (r *repositoryProductImpl) CreateProduct(product models.Product) (models.Product, error) {
	if err := r.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

// DeleteProduct implements RepositoryProduct.
func (r *repositoryProductImpl) DeleteProduct(product models.Product) (models.Product, error) {
	if err := r.DB.Delete(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// FindAllProduct implements RepositoryProduct.
func (r *repositoryProductImpl) FindAllProduct() ([]models.Product, error) {
	var product []models.Product

	if err := r.DB.Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// FindProduct implements RepositoryProduct.
func (r *repositoryProductImpl) FindProduct(id int) (models.Product, error) {
	var product models.Product

	if err := r.DB.Find(&product, id).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// UpdateProduct implements RepositoryProduct.
func (r *repositoryProductImpl) UpdateProduct(product models.Product) (models.Product, error) {
	if err := r.DB.Save(&product).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

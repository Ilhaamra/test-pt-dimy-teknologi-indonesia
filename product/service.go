package product

import (
	"test-dimy-teknologi/models"
)

type ServiceProduct interface {
	FindAllProduct() ([]models.Product, error)
	FindProduct(id int) (models.Product, error)
	CreateProduct(Product RequestProduct) (models.Product, error)
	UpdateProduct(id int, product RequestProduct) (models.Product, error)
	DeleteProduct(id int) (models.Product, error)
}

type serviceProductImpl struct {
	RepositoryProduct RepositoryProduct
}

func NewServiceProduct(repositoryProduct RepositoryProduct) ServiceProduct {
	return &serviceProductImpl{
		RepositoryProduct: repositoryProduct,
	}
}

// CreateProduct implements ServiceProduct.
func (s *serviceProductImpl) CreateProduct(Product RequestProduct) (models.Product, error) {
	requestProduct := models.Product{
		Name:  Product.Name,
		Price: Product.Price,
	}
	newProduct, err := s.RepositoryProduct.CreateProduct(requestProduct)
	if err != nil {
		return models.Product{}, err
	}
	return newProduct, nil
}

// DeleteProduct implements ServiceProduct.
func (s *serviceProductImpl) DeleteProduct(id int) (models.Product, error) {
	p, err := s.RepositoryProduct.FindProduct(id)
	if err != nil {
		panic(err)
	}
	delProduct, err := s.RepositoryProduct.DeleteProduct(p)

	return delProduct, err
}

// FindAllProduct implements ServiceProduct.
func (s *serviceProductImpl) FindAllProduct() ([]models.Product, error) {
	products, err := s.RepositoryProduct.FindAllProduct()
	return products, err
}

// FindProduct implements ServiceProduct.
func (s *serviceProductImpl) FindProduct(id int) (models.Product, error) {
	product, err := s.RepositoryProduct.FindProduct(id)
	return product, err
}

// UpdateProduct implements ServiceProduct.
func (s *serviceProductImpl) UpdateProduct(id int, product RequestProduct) (models.Product, error) {
	p, err := s.RepositoryProduct.FindProduct(id)
	if err != nil {
		panic(err)
	}

	p.Name = product.Name
	p.Price = product.Price

	newProduct, err := s.RepositoryProduct.UpdateProduct(p)

	return newProduct, err
}

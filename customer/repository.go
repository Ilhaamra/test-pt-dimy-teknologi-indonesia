package customer

import (
	"test-dimy-teknologi/models"

	"gorm.io/gorm"
)

type RepositoryCustomer interface {
	FindAllCustomer() ([]models.Customer, error)
	FindCustomer(id int) (models.Customer, error)
	CreateCustomer(customer models.Customer) (models.Customer, error)
	UpdateCustomer(customer models.Customer) (models.Customer, error)
	DeleteCustomer(customer models.Customer) (models.Customer, error)
	SaveAddress(address models.CustomerAddress) (models.CustomerAddress, error)
	FindAddress(id int) (models.CustomerAddress, error)
}

type repositoryCustomerImpl struct {
	DB *gorm.DB
}

func NewRespositoryCustomer(db *gorm.DB) RepositoryCustomer {
	return &repositoryCustomerImpl{
		DB: db,
	}
}

// CreateCustomer implements RepositoryCustomer.
func (r *repositoryCustomerImpl) CreateCustomer(customer models.Customer) (models.Customer, error) {
	if err := r.DB.Create(&customer).Error; err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

// UpdateCustomer implements RepositoryCustomer.
func (r *repositoryCustomerImpl) UpdateCustomer(customer models.Customer) (models.Customer, error) {
	if err := r.DB.Save(&customer).Error; err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

// FindAllCustomer implements RepositoryCustomer.
func (r *repositoryCustomerImpl) FindAllCustomer() ([]models.Customer, error) {
	var customer []models.Customer

	if err := r.DB.Find(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

// FindCustomer implements RepositoryCustomer.
func (r *repositoryCustomerImpl) FindCustomer(id int) (models.Customer, error) {
	var customer models.Customer

	if err := r.DB.Find(&customer, id).Error; err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

// DeleteCustomer implements RepositoryCustomer.
func (r *repositoryCustomerImpl) DeleteCustomer(customer models.Customer) (models.Customer, error) {
	if err := r.DB.Delete(&customer).Error; err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

// FindAddress implements RepositoryCustomer.
func (r *repositoryCustomerImpl) FindAddress(id int) (models.CustomerAddress, error) {
	var address models.CustomerAddress

	if err := r.DB.Find(&address, id).Error; err != nil {
		return models.CustomerAddress{}, err
	}
	return address, nil
}

// SaveAddress implements RepositoryCustomer.
func (r *repositoryCustomerImpl) SaveAddress(address models.CustomerAddress) (models.CustomerAddress, error) {
	if err := r.DB.Create(&address).Error; err != nil {
		return models.CustomerAddress{}, err
	}

	return address, nil
}

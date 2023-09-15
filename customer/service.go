package customer

import (
	"test-dimy-teknologi/models"
)

type ServiceCustomer interface {
	FindAllCustomer() ([]models.Customer, error)
	FindCustomer(id int) (models.Customer, error)
	CreateCustomer(customer RequestCustomer) (models.Customer, error)
	UpdateCustomer(id int, customer RequestCustomer) (models.Customer, error)
	DeleteCustomer(id int) (models.Customer, error)
	FindAddress(id int) (models.CustomerAddress, error)
	SaveAddress(address RequestCustomerAddress) (models.CustomerAddress, error)
}

type serviceCustomerImpl struct {
	RepositoryCustomer RepositoryCustomer
}

func NewServiceCustomer(repositoryCustomer RepositoryCustomer) ServiceCustomer {
	return &serviceCustomerImpl{
		RepositoryCustomer: repositoryCustomer,
	}
}

// CreateCustomer implements ServiceCustomer.
func (s *serviceCustomerImpl) CreateCustomer(customer RequestCustomer) (models.Customer, error) {
	requestCustomer := models.Customer{
		Name: customer.Name,
	}
	newCustomer, err := s.RepositoryCustomer.CreateCustomer(requestCustomer)
	if err != nil {
		return models.Customer{}, err
	}
	return newCustomer, nil
}

// UpdateCustomer implements ServiceCustomer.
func (s *serviceCustomerImpl) UpdateCustomer(id int, customer RequestCustomer) (models.Customer, error) {
	c, err := s.RepositoryCustomer.FindCustomer(id)
	if err != nil {
		panic(err)
	}

	c.Name = customer.Name

	newCustomer, err := s.RepositoryCustomer.UpdateCustomer(c)

	return newCustomer, err
}

// FindAllCustomer implements ServiceCustomer.
func (s *serviceCustomerImpl) FindAllCustomer() ([]models.Customer, error) {
	customers, err := s.RepositoryCustomer.FindAllCustomer()
	return customers, err
}

// FindCustomer implements ServiceCustomer.
func (s *serviceCustomerImpl) FindCustomer(id int) (models.Customer, error) {
	customer, err := s.RepositoryCustomer.FindCustomer(id)
	return customer, err
}

// DeleteCustomer implements ServiceCustomer.
func (s *serviceCustomerImpl) DeleteCustomer(id int) (models.Customer, error) {
	c, err := s.RepositoryCustomer.FindCustomer(id)
	if err != nil {
		panic(err)
	}
	delCustomer, err := s.RepositoryCustomer.DeleteCustomer(c)

	return delCustomer, err
}

// FindAddress implements ServiceCustomer.
func (s *serviceCustomerImpl) FindAddress(id int) (models.CustomerAddress, error) {
	address, err := s.RepositoryCustomer.FindAddress(id)
	return address, err
}

// SaveAddress implements ServiceCustomer.
func (s *serviceCustomerImpl) SaveAddress(address RequestCustomerAddress) (models.CustomerAddress, error) {

	customer, err := s.FindCustomer(address.CustomerId)
	if err != nil {
		return models.CustomerAddress{}, err
	}

	requestAddress := models.CustomerAddress{
		CustomerId:  customer.Id,
		AddressDate: address.AddressDate,
	}
	newAddress, err := s.RepositoryCustomer.SaveAddress(requestAddress)
	if err != nil {
		return models.CustomerAddress{}, err
	}
	return newAddress, nil
}

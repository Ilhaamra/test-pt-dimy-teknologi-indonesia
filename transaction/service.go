package transaction

import (
	"test-dimy-teknologi/customer"
	"test-dimy-teknologi/models"
)

type ServiceTransaction interface {
	FindAllTransaction() ([]models.Transaction, error)
	FindTransaction(id int) (models.Transaction, error)
	CreateTransaction(transaction RequestTransaction) (models.Transaction, error)
	SaveTransactionProduct(tp RequestTransactionProduct) (models.TransactionProduct, error)
	GetTransactionProduct() ([]models.TransactionProduct, error)
}

type serviceTransactionImpl struct {
	RepositoryTransaction RepositoryTransaction
	Customer              customer.RepositoryCustomer
}

func NewServiceTransaction(repositoryTransaction RepositoryTransaction, customer customer.RepositoryCustomer) ServiceTransaction {
	return &serviceTransactionImpl{
		RepositoryTransaction: repositoryTransaction,
		Customer:              customer,
	}
}

// CreateTransaction implements ServiceTransaction.
func (s *serviceTransactionImpl) CreateTransaction(transaction RequestTransaction) (models.Transaction, error) {
	customer, err := s.Customer.FindCustomer(transaction.CustomerId)
	if err != nil {
		return models.Transaction{}, err
	}
	address, err := s.Customer.FindAddress(transaction.CustomerAddressId)
	if err != nil {
		return models.Transaction{}, err
	}

	request := models.Transaction{
		CustomerId:        customer.Id,
		CustomerAddressId: address.Id,
		TransactionDate:   transaction.TransactionDate,
	}
	newTransaction, err := s.RepositoryTransaction.CreateTransaction(request)
	if err != nil {
		return models.Transaction{}, err
	}
	return newTransaction, nil
}

// FindAllTransaction implements ServiceTransaction.
func (s *serviceTransactionImpl) FindAllTransaction() ([]models.Transaction, error) {
	transactions, err := s.RepositoryTransaction.FindAllTransaction()
	return transactions, err
}

// FindTransaction implements ServiceTransaction.
func (s *serviceTransactionImpl) FindTransaction(id int) (models.Transaction, error) {
	transactions, err := s.RepositoryTransaction.FindTransaction(id)
	return transactions, err
}

// GetTransactionProduct implements ServiceTransaction.
func (*serviceTransactionImpl) GetTransactionProduct() ([]models.TransactionProduct, error) {
	panic("unimplemented")
}

// SaveTransactionProduct implements ServiceTransaction.
func (*serviceTransactionImpl) SaveTransactionProduct(tp RequestTransactionProduct) (models.TransactionProduct, error) {
	panic("unimplemented")
}

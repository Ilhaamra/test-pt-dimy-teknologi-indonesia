package transaction

import (
	"test-dimy-teknologi/models"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	FindAllTransaction() ([]models.Transaction, error)
	FindTransaction(id int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	SaveTransactionProduct(tp models.TransactionProduct) (models.TransactionProduct, error)
	GetTransactionProduct() ([]models.TransactionProduct, error)
}

type repositoryTransactionImpl struct {
	DB *gorm.DB
}

func NewRespositoryTransaction(db *gorm.DB) RepositoryTransaction {
	return &repositoryTransactionImpl{
		DB: db,
	}
}

// CreateTransaction implements RepositoryTransaction.
func (r *repositoryTransactionImpl) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := r.DB.Create(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

// FindAllTransaction implements RepositoryTransaction.
func (r *repositoryTransactionImpl) FindAllTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction

	if err := r.DB.Table("transaction").Select("id AS transaction_id, customer_id, customer_address_id, transaction_date").Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// FindTransaction implements RepositoryTransaction.
func (r *repositoryTransactionImpl) FindTransaction(id int) (models.Transaction, error) {
	var transaction models.Transaction

	if err := r.DB.Table("transaction").Select("id AS transaction_id, customer_id, customer_address_id, transaction_date").Where("customer_id = ?", id).Find(&transaction).Error; err != nil {
		return models.Transaction{}, err
	}
	return transaction, nil
}

// SaveTransactionProduct implements RepositoryTransaction.
func (r *repositoryTransactionImpl) SaveTransactionProduct(tp models.TransactionProduct) (models.TransactionProduct, error) {
	if err := r.DB.Create(&tp).Error; err != nil {
		return models.TransactionProduct{}, err
	}

	return tp, nil
}

// GetTransaction implements RepositoryTransaction.
func (r *repositoryTransactionImpl) GetTransactionProduct() ([]models.TransactionProduct, error) {
	var transaction []models.TransactionProduct

	if err := r.DB.Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

package controllers

import (
	"fmt"
	"net/http"
	"test-dimy-teknologi/models"
	"test-dimy-teknologi/transaction"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ControllerTransaction interface {
	FindAllTransaction(ctx *gin.Context)
	FindTransaction(ctx *gin.Context)
	CreateTransaction(ctx *gin.Context)
}

type controllerTransactionImpl struct {
	ServiceTransaction transaction.ServiceTransaction
}

func NewControllerTransaction(serviceTransaction transaction.ServiceTransaction) ControllerTransaction {
	return &controllerTransactionImpl{
		ServiceTransaction: serviceTransaction,
	}
}

// CreateTransaction implements ControllerTransaction.
func (c *controllerTransactionImpl) CreateTransaction(ctx *gin.Context) {
	var transactionRequest transaction.RequestTransaction

	err := ctx.ShouldBindJSON(&transactionRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	transaction, err := c.ServiceTransaction.CreateTransaction(transactionRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": toTransactionResponse(transaction),
	})
}

// FindAllTransaction implements ControllerTransaction.
func (c *controllerTransactionImpl) FindAllTransaction(ctx *gin.Context) {
	transactions, err := c.ServiceTransaction.FindAllTransaction()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var response []transaction.ResponseTransaction
	for _, t := range transactions {
		responseTransaction := toTransactionResponse(t)
		response = append(response, responseTransaction)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// FindTransaction implements ControllerTransaction.
func (c *controllerTransactionImpl) FindTransaction(ctx *gin.Context) {
	panic("unimplemented")
}

func toTransactionResponse(p models.Transaction) transaction.ResponseTransaction {
	return transaction.ResponseTransaction{
		CustomerId:        p.CustomerId,
		CustomerAddressId: p.CustomerAddressId,
		TransactionDate:   p.TransactionDate,
	}
}

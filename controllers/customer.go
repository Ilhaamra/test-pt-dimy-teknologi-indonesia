package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"test-dimy-teknologi/customer"
	"test-dimy-teknologi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ControllerCustomer interface {
	FindAllCustomer(ctx *gin.Context)
	FindCustomer(ctx *gin.Context)
	CreateCustomer(ctx *gin.Context)
	UpdateCustomer(ctx *gin.Context)
	DeleteCustomer(ctx *gin.Context)
	FindAddress(ctx *gin.Context)
	SaveAddress(ctx *gin.Context)
}

type controllerCustomerImpl struct {
	ServiceCustomer customer.ServiceCustomer
}

func NewControllerCustomer(serviceCustomer customer.ServiceCustomer) ControllerCustomer {
	return &controllerCustomerImpl{
		ServiceCustomer: serviceCustomer,
	}
}

// CreateCustomer implements ControllerCustomer.
func (c *controllerCustomerImpl) CreateCustomer(ctx *gin.Context) {
	var customerRequest customer.RequestCustomer

	err := ctx.ShouldBindJSON(&customerRequest)
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
	customer, err := c.ServiceCustomer.CreateCustomer(customerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": toCustomerResponse(customer),
	})
}

// UpdateCustomer implements ControllerCustomer.
func (c *controllerCustomerImpl) UpdateCustomer(ctx *gin.Context) {
	var customerRequest customer.RequestCustomer

	err := ctx.ShouldBindJSON(&customerRequest)
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
	id, _ := strconv.Atoi(ctx.Param("id"))
	customer, err := c.ServiceCustomer.UpdateCustomer(id, customerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": toCustomerResponse(customer),
	})
}

// FindAllCustomer implements ControllerCustomer.
func (c *controllerCustomerImpl) FindAllCustomer(ctx *gin.Context) {
	customers, err := c.ServiceCustomer.FindAllCustomer()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var response []customer.ResponseCustomer
	for _, c := range customers {
		responseCustomer := toCustomerResponse(c)
		response = append(response, responseCustomer)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// FindCustomer implements ControllerCustomer.
func (c *controllerCustomerImpl) FindCustomer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	customer, err := c.ServiceCustomer.FindCustomer(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	response := toCustomerResponse(customer)

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// DeleteCustomer implements ControllerCustomer.
func (c *controllerCustomerImpl) DeleteCustomer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	customer, err := c.ServiceCustomer.DeleteCustomer(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	response := toCustomerResponse(customer)

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// FindAddress implements ControllerCustomer.
func (c *controllerCustomerImpl) FindAddress(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	customer, err := c.ServiceCustomer.FindAddress(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	response := toAddressResponse(customer)

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// SaveAddress implements ControllerCustomer.
func (c *controllerCustomerImpl) SaveAddress(ctx *gin.Context) {
	var customerRequest customer.RequestCustomerAddress

	err := ctx.ShouldBindJSON(&customerRequest)
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
	customer, err := c.ServiceCustomer.SaveAddress(customerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": toAddressResponse(customer),
	})
}

func toAddressResponse(c models.CustomerAddress) customer.ResponseCustomerAddress {
	return customer.ResponseCustomerAddress{
		Id:          c.Id,
		CustomerId:  c.CustomerId,
		AddressDate: c.AddressDate,
	}
}
func toCustomerResponse(c models.Customer) customer.ResponseCustomer {
	return customer.ResponseCustomer{
		Id:   c.Id,
		Name: c.Name,
	}
}

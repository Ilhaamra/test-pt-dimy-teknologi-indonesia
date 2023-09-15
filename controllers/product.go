package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"test-dimy-teknologi/models"
	"test-dimy-teknologi/product"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ControllerProduct interface {
	FindAllProduct(ctx *gin.Context)
	FindProduct(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
}

type controllerProductImpl struct {
	ServiceProduct product.ServiceProduct
}

func NewControllerProduct(serviceProduct product.ServiceProduct) ControllerProduct {
	return &controllerProductImpl{
		ServiceProduct: serviceProduct,
	}
}

// CreateProduct implements ControllerProduct.
func (c *controllerProductImpl) CreateProduct(ctx *gin.Context) {
	var productRequest product.RequestProduct

	err := ctx.ShouldBindJSON(&productRequest)
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
	product, err := c.ServiceProduct.CreateProduct(productRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": toProductResponse(product),
	})
}

// FindAllProduct implements ControllerProduct.
func (c *controllerProductImpl) FindAllProduct(ctx *gin.Context) {
	products, err := c.ServiceProduct.FindAllProduct()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var response []product.ResponseProduct
	for _, p := range products {
		responseProduct := toProductResponse(p)
		response = append(response, responseProduct)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// FindProduct implements ControllerProduct.
func (c *controllerProductImpl) FindProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	product, err := c.ServiceProduct.FindProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	response := toProductResponse(product)

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func toProductResponse(p models.Product) product.ResponseProduct {
	return product.ResponseProduct{
		Id:    p.Id,
		Name:  p.Name,
		Price: p.Price,
	}
}

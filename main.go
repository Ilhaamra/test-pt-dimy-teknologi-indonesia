package main

import (
	"test-dimy-teknologi/config"
	"test-dimy-teknologi/controllers"
	"test-dimy-teknologi/customer"
	"test-dimy-teknologi/product"
	"test-dimy-teknologi/transaction"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := config.Connction()

	customerRepository := customer.NewRespositoryCustomer(db)
	customerServie := customer.NewServiceCustomer(customerRepository)
	customerController := controllers.NewControllerCustomer(customerServie)

	productRepository := product.NewRespositoryProduct(db)
	productService := product.NewServiceProduct(productRepository)
	productController := controllers.NewControllerProduct(productService)

	transactionRepository := transaction.NewRespositoryTransaction(db)
	transactionService := transaction.NewServiceTransaction(transactionRepository, customerRepository)
	transactionController := controllers.NewControllerTransaction(transactionService)

	v1 := router.Group("v1")

	v1.POST("/customer", customerController.CreateCustomer)
	v1.POST("/customer/address", customerController.SaveAddress)
	v1.GET("/customer/address/:id", customerController.FindAddress)
	v1.GET("/customers", customerController.FindAllCustomer)
	v1.GET("/customer/:id", customerController.FindCustomer)
	v1.PUT("/customer/:id", customerController.UpdateCustomer)
	v1.DELETE("/customer/:id", customerController.DeleteCustomer)

	v1.POST("/product", productController.CreateProduct)
	v1.GET("/products", productController.FindAllProduct)
	v1.GET("/product/:id", productController.FindProduct)

	v1.POST("/transaction", transactionController.CreateTransaction)
	v1.GET("/transactions", transactionController.FindAllTransaction)

	router.Run(":9000")
}

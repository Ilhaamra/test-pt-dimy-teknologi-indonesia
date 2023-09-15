# Test PT.DIMY Teknologi Indonesia

## How to run program?

```
go mod tity
go run main.go
```

## API Endpont

```
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
```

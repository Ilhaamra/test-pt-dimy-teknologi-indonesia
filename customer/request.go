package customer

type RequestCustomer struct {
	Name string `json:"name" binding:"required"`
}

type RequestCustomerAddress struct {
	CustomerId  int    `json:"customer_id" binding:"required"`
	AddressDate string `json:"address_date" binding:"required"`
}

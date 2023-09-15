package transaction

type RequestTransaction struct {
	CustomerId        int    `json:"customer_id" binding:"required,number"`
	CustomerAddressId int    `json:"customer_address_id" binding:"required,number"`
	TransactionDate   string `json:"transaction_date"`
}

type RequestTransactionProduct struct {
	TransactionId int `json:"transaction_id" binding:"required,number"`
	ProductId     int `json:"product_id" binding:"required,number"`
	Quantity      int `json:"quantity" binding:"required,number"`
}

type RequestTransactionPaymentMethod struct {
	TransactionId   int `json:"transaction_id" binding:"required,number"`
	PaymentMethodId int `json:"payment_method_id" binding:"required,number"`
}

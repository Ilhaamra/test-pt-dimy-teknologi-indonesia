package transaction

type ResponseTransaction struct {
	CustomerId        int    `json:"customer_id"`
	CustomerAddressId int    `json:"customer_address_id"`
	TransactionDate   string `json:"transaction_date"`
}

type ResponseTransactionProduct struct {
	TransactionId int `json:"transaction_id"`
	ProductId     int `json:"product_id"`
	Quantity      int `json:"quantity"`
}

type ResponseTransactionPaymentMethod struct {
	TransactionId   int `json:"transaction_id"`
	PaymentMethodId int `json:"payment_method_id"`
}

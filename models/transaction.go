package models

type Transaction struct {
	Id                int    `gorm:"column:id;primaryKey"`
	CustomerId        int    `gorm:"column:customer_id"`
	CustomerAddressId int    `gorm:"column:customer_address_id"`
	TransactionDate   string `gorm:"column:transaction_date"`
}

type TransactionProduct struct {
	Id            int
	TransactionId int
	ProductId     int
	Quantity      int
}

type TransactionPaymentMethod struct {
	Id              int
	TransactionId   int
	PaymentMethodId int
}

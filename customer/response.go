package customer

type ResponseCustomer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type ResponseCustomerAddress struct {
	Id          int    `json:"id"`
	CustomerId  int    `json:"customer_id"`
	AddressDate string `json:"address_date"`
}

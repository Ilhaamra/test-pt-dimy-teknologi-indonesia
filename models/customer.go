package models

type Customer struct {
	Id   int
	Name string
}

type CustomerAddress struct {
	Id          int
	CustomerId  int
	AddressDate string
}

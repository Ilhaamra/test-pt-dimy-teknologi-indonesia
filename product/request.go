package product

type RequestProduct struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required,number"`
}

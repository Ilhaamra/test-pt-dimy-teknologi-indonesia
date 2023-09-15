package payment

type RequestPayment struct {
	Name     string `json:"name" binding:"required"`
	IsActive int    `json:"is_active" binding:"required"`
}

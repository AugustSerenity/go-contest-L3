package dto

type ItemRequest struct {
	Type     string  `json:"type" binding:"required"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount" binding:"required,gte=0"`
	Date     string  `json:"date" binding:"required"`
}

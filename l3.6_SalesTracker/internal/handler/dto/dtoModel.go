package dto

type ItemRequest struct {
	Type     string  `json:"type" binding:"required"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount" binding:"required,gte=0"`
	Date     string  `json:"date" binding:"required"`
}

type AnalyticsRequest struct {
	From     *string `form:"from"`
	To       *string `form:"to"`
	Category *string `form:"category"`
}

type AnalyticsResponse struct {
	Sum    float64 `json:"sum"`
	Avg    float64 `json:"avg"`
	Count  int64   `json:"count"`
	Median float64 `json:"median"`
	P90    float64 `json:"p90"`
}

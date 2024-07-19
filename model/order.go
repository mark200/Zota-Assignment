package model

type Order struct {
	OrderId       string  `json:"orderId"`
	OrderAmount   float64 `json:"orderAmount"`
	OrderCurrency string  `json:"orderCurrency"`
}

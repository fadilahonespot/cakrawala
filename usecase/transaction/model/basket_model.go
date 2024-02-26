package model

type BasketResponse struct {
	ProductID   int     `json:"productId"`
	Name        string  `json:"name"`
	Weight      int     `json:"weight"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Qty         int     `json:"quantity"`
	TotalPrice  float64 `json:"totalPrice"`
	TotalWeight int     `json:"totalWeight"`
}

type AddBasketRequest struct {
	ProductID int `json:"productId" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

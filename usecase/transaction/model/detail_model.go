package model

import "time"

type DetailTransactionResponse struct {
	TransactionId   string    `json:"transactionId"`
	UserName        string    `json:"userName"`
	UserAddress     string    `json:"userAddress"`
	UserPhone       string    `json:"userPhone"`
	UserEmail       string    `json:"userEmail"`
	Xpayment        string    `json:"xPayment"`
	Status          string    `json:"status"`
	Bank            string    `json:"bank"`
	VirtualAccount  string    `json:"virtualAccount"`
	AdminFree       int       `json:"adminFree"`
	ShippingCost    int       `json:"shippingCost"`
	ProductAmount   int       `json:"productAmount"`
	TotalPayment    int       `json:"totalPayment"`
	ProductQuantity int       `json:"productQuantity"`
	ProductWeight   int       `json:"productWeight"`
	CourierName     string    `json:"courierName"`
	CourierService  string    `json:"courierService"`
	OriginCity      string    `json:"originCity"`
	DestinationCity string    `json:"destinationCity"`
	Etd             string    `json:"etd"`
	CreatedAt       time.Time `json:"createdAt"`
	ExpiredAt       time.Time `json:"expiredAt"`
	Items           []TrxItem `json:"items"`
}

type TrxItem struct {
	ProductID   int     `json:"productId"`
	ProductName string  `json:"productName"`
	ProductImg  string  `json:"productImg"`
	Price       float64 `json:"price"`
	Weight      int     `json:"weight"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"totalPrice"`
	TotalWeight float64 `json:"totalWeight"`
}

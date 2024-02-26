package model

import "time"

type CheckoutRequest struct {
	BankCode       string `json:"bankCode"`
	CourierCode    string `json:"courierCode"`
	CourierService string `json:"courierService"`
}

type CheckoutResponse struct {
	TransactionID      string    `json:"transactionId"`
	XPayment           string    `json:"xPayment"`
	BankCode           string    `json:"bankCode"`
	VirtualAccount     string    `json:"virtualAccount"`
	VirtualAccountName string    `json:"virtualAccountName"`
	ShippingCost       int       `json:"shippingCost"`
	ProductPrice       int       `json:"productPrice"`
	AdminFree          int       `json:"adminFree"`
	TotalPayment       int       `json:"totalAmount"`
	ExpiredDate        time.Time `json:"expiredDate"`
}

package model

import "time"

type HistoryTransactionResponse struct {
	TransactionId string    `json:"transactionId"`
	Xpayment      string    `json:"xPayment"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	Amount        int       `json:"amount"`
}

package xendit

import "time"

type GetBankResponse struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	IsActivated bool   `json:"is_activated"`
}

type CreateVirtualAccountRequest struct {
	ExternalID     string    `json:"external_id"`
	BankCode       string    `json:"bank_code"`
	Name           string    `json:"name"`
	IsSingleUse    bool      `json:"is_single_use"`
	IsClosed       bool      `json:"is_closed"`
	ExpectedAmount int       `json:"expected_amount"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type CreateViartualAccountResponse struct {
	ID             string    `json:"id"`
	OwnerID        string    `json:"owner_id"`
	ExternalID     string    `json:"external_id"`
	AccountNumber  string    `json:"account_number"`
	BankCode       string    `json:"bank_code"`
	MerchantCode   string    `json:"merchant_code"`
	Name           string    `json:"name"`
	IsClosed       bool      `json:"is_closed"`
	ExpectedAmount int       `json:"expected_amount"`
	ExpirationDate time.Time `json:"expiration_date"`
	IsSingleUse    bool      `json:"is_single_use"`
	Status         string    `json:"status"`
	Currency       string    `json:"currency"`
	Country        string    `json:"country"`
}

type CheckPaymentResponse struct {
	ID                       string        `json:"id"`
	PaymentID                string        `json:"payment_id"`
	CallbackVirtualAccountID string        `json:"callback_virtual_account_id"`
	OwnerID                  string        `json:"owner_id"`
	ExternalID               string        `json:"external_id"`
	BankCode                 string        `json:"bank_code"`
	MerchantCode             string        `json:"merchant_code"`
	AccountNumber            string        `json:"account_number"`
	Amount                   int           `json:"amount"`
	Currency                 string        `json:"currency"`
	TransactionTimestamp     time.Time     `json:"transaction_timestamp"`
	SenderName               string        `json:"sender_name"`
	PaymentDetail            PaymentDetail `json:"payment_detail"`
}

type PaymentDetail struct {
	PaymentInterface    string `json:"payment_interface"`
	Remark              string `json:"remark"`
	Reference           string `json:"reference"`
	SenderAccountNumber string `json:"sender_account_number"`
	SenderChannelCode   string `json:"sender_channel_code"`
	SenderName          string `json:"sender_name"`
	TransferMethod      string `json:"transfer_method"`
}

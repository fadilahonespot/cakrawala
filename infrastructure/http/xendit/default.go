package xendit

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/library/rest"
)

type wrapper struct {
	client rest.RestClient
}

var address string

func NewWrapper() XenditWrapper {
	restOptions := rest.Options{
		Address: os.Getenv("XENDIT_HOST"),
		Timeout: time.Duration(10 * time.Second),
		SkipTLS: false,
	}
	client := rest.New(restOptions)
	address = restOptions.Address

	return &wrapper{client: client}
}

func getRequestHeaders(ctx context.Context) (headers http.Header) {
	token := os.Getenv("XENDIT_TOKEN") + ":"
	auth := fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(token)))

	headers = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{auth},
	}

	return
}

func (w *wrapper) GetBank(ctx context.Context) (resp []GetBankResponse, err error) {
	path := "/available_virtual_account_banks"

	logger.Info(ctx, "[GetBank Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[Xendit] GetBank error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Xendit] GetBank return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Xendit] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[GetBank Response]", address+path, resp)

	return
}

func (w *wrapper) CreateVirtualAccount(ctx context.Context, req CreateVirtualAccountRequest) (resp CreateViartualAccountResponse, err error) {
	path := "/callback_virtual_accounts"

	logger.Info(ctx, "[CreateVirtualAccount Request]", address+path, req)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Post(ctx, path, headers, req)
	if err != nil {
		err = fmt.Errorf("[Xendit] CreateVirtualAccount error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Xendit] CreateVirtualAccount return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Xendit] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[CreateVirtualAccount Response]", address+path, resp)

	return
}

func (w *wrapper) CheckVirtualAccount(ctx context.Context, id string) (resp CreateViartualAccountResponse, err error) {
	path := fmt.Sprintf("/callback_virtual_accounts/%v", id)

	logger.Info(ctx, "[CheckVirtualAccount Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[Xendit] CheckVirtualAccount error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Xendit] CheckVirtualAccount return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Xendit] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[CheckVirtualAccount Response]", address+path, resp)

	return
}

func (w *wrapper) CheckPayment(ctx context.Context, paymentId string) (resp CheckPaymentResponse, err error) {
	path := fmt.Sprintf("/callback_virtual_account_payments/payment_id=%v", paymentId)

	logger.Info(ctx, "[CheckPayment Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[Xendit] CheckPayment error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Xendit] CheckPayment return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Xendit] Unmarshal Response Error %v", err.Error())
	}

	logger.Info(ctx, "[CheckPayment Response]", address+path, resp)

	return
}

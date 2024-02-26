package transaction

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/domain/repository"
	"github.com/fadilahonespot/cakrawala/infrastructure/cached"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/mailjet"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/rajaongkir"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/xendit"
	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
)

type defaultTransaction struct {
	rajaongkirWrapper rajaongkir.RajaOngkirWrapper
	userRepo          repository.UserRepository
	productRepo       repository.ProductRepository
	xenditWrapper     xendit.XenditWrapper
	courierInfoRepo   repository.CourierInfoRepository
	transactionRepo   repository.TransactionRepository
	courierRepo       repository.CourierRepository
	paymentInfoRepo   repository.PaymentInfoRepository
	mailjetWrapper    mailjet.MailjetWrapper
	cache             cached.RedisClient
	addressRepo       repository.AddressRepository
}

func SetupTransactionService() *defaultTransaction {
	return &defaultTransaction{}
}

func (s *defaultTransaction) SetRajaongkirWrapper(wrapper rajaongkir.RajaOngkirWrapper) *defaultTransaction {
	s.rajaongkirWrapper = wrapper
	return s
}

func (s *defaultTransaction) SetUserRepo(repo repository.UserRepository) *defaultTransaction {
	s.userRepo = repo
	return s
}

func (s *defaultTransaction) SetProductRepo(repo repository.ProductRepository) *defaultTransaction {
	s.productRepo = repo
	return s
}

func (s *defaultTransaction) SetXenditWrapper(wrapper xendit.XenditWrapper) *defaultTransaction {
	s.xenditWrapper = wrapper
	return s
}

func (s *defaultTransaction) SetCourierInfoRepo(repo repository.CourierInfoRepository) *defaultTransaction {
	s.courierInfoRepo = repo
	return s
}

func (s *defaultTransaction) SetTransactionRepo(repo repository.TransactionRepository) *defaultTransaction {
	s.transactionRepo = repo
	return s
}

func (s *defaultTransaction) SetCourierRepo(repo repository.CourierRepository) *defaultTransaction {
	s.courierRepo = repo
	return s
}

func (s *defaultTransaction) SetPaymentInfoRepo(repo repository.PaymentInfoRepository) *defaultTransaction {
	s.paymentInfoRepo = repo
	return s
}

func (s *defaultTransaction) SetMailjetWrapper(wrapper mailjet.MailjetWrapper) *defaultTransaction {
	s.mailjetWrapper = wrapper
	return s
}

func (s *defaultTransaction) SetCache(cache cached.RedisClient) *defaultTransaction {
	s.cache = cache
	return s
}

func (s *defaultTransaction) SetAddressRepo(repo repository.AddressRepository) *defaultTransaction {
	s.addressRepo = repo
	return s
}

func (s *defaultTransaction) Validate() TransactionService {
	if s.rajaongkirWrapper == nil {
		panic("rajaongkir wrapper is nil")
	}

	if s.userRepo == nil {
		panic("user repo is nil")
	}

	if s.productRepo == nil {
		panic("product repo is nil")
	}

	if s.xenditWrapper == nil {
		panic("xendit warapper is nil")
	}

	if s.courierInfoRepo == nil {
		panic("courier info repo is nil")
	}

	if s.transactionRepo == nil {
		panic("transaction repo is nil")
	}

	if s.courierRepo == nil {
		panic("courier repo is nil")
	}

	if s.paymentInfoRepo == nil {
		panic("payment info repo is nil")
	}

	if s.mailjetWrapper == nil {
		panic("mailjet wrapper is nil")
	}

	if s.cache == nil {
		panic("cache is nil")
	}

	if s.addressRepo == nil {
		panic("address repo is nil")
	}

	return s
}

func generatePaymentNotifBody(req model.CheckoutResponse) (emailBody string, err error) {
	local, _ := time.LoadLocation("Asia/Jakarta")
	configPath := "./resources/template/payment_notification.html"
	tempByte, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	emailBody = string(tempByte)
	emailBody = strings.ReplaceAll(emailBody, "[Nomor Pesanan]", req.XPayment)
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Transaksi]", time.Now().In(local).Format("02 January 2006 15:04"))
	emailBody = strings.ReplaceAll(emailBody, "[Ongkos Kirim]", strconv.Itoa(req.ShippingCost))
	emailBody = strings.ReplaceAll(emailBody, "[Harga Produk]", strconv.Itoa(req.ProductPrice))
	emailBody = strings.ReplaceAll(emailBody, "[Total Pembayaran]", strconv.Itoa(req.TotalPayment))
	emailBody = strings.ReplaceAll(emailBody, "[Metode Pembayaran]", "Virtual Account")
	emailBody = strings.ReplaceAll(emailBody, "[Nomor Virtual Account]", req.VirtualAccount)
	emailBody = strings.ReplaceAll(emailBody, "[Nama Bank]", req.BankCode)
	emailBody = strings.ReplaceAll(emailBody, "[Waktu Kadaluarsa]", req.ExpiredDate.In(local).Format("02 January 2006 15:04"))
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Notifikasi]", time.Now().In(local).Format("02 January 2006"))

	return
}

func generatePaymentSuccessBody(req xendit.CheckPaymentResponse, paymentData entity.PaymentInfo) (emailBody string, err error) {
	local, _ := time.LoadLocation("Asia/Jakarta")
	configPath := "./resources/template/payment_notif_success.html"
	tempByte, err := os.ReadFile(configPath)
	if err != nil {
		return
	}
	emailBody = string(tempByte)
	emailBody = strings.ReplaceAll(emailBody, "[Nomor Pesanan]", req.ExternalID)
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Transaksi]", paymentData.CreatedAt.In(local).Format("02 January 2006 15:04"))
	emailBody = strings.ReplaceAll(emailBody, "[Total Pembayaran]", strconv.Itoa(req.Amount))
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Pembayaran]", req.TransactionTimestamp.In(local).Format("02 January 2006 15:04"))
	emailBody = strings.ReplaceAll(emailBody, "[Nomor Virtual Account]", paymentData.AccountNumber)
	emailBody = strings.ReplaceAll(emailBody, "[Nama Bank]", req.BankCode)
	emailBody = strings.ReplaceAll(emailBody, "[Tanggal Notifikasi]", time.Now().In(local).Format("02 January 2006"))

	return
}

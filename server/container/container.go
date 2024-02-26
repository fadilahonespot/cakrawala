package container

import (
	"fmt"
	"os"

	"github.com/fadilahonespot/cakrawala/infrastructure/cached"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/filebox"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/mailjet"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/rajaongkir"
	"github.com/fadilahonespot/cakrawala/infrastructure/http/xendit"
	"github.com/fadilahonespot/cakrawala/infrastructure/mysql"
	"github.com/fadilahonespot/cakrawala/server/handler"
	"github.com/fadilahonespot/cakrawala/server/router"
	"github.com/fadilahonespot/cakrawala/usecase/product"
	"github.com/fadilahonespot/cakrawala/usecase/transaction"
	"github.com/fadilahonespot/cakrawala/usecase/user"
	"github.com/fadilahonespot/cakrawala/utils/database"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/labstack/echo/v4"
)

func NewContainer() {
	// Setup Logger
	logger.NewLogger()

	// Setup Database
	db := database.InitDB()

	// Setup Redis
	cache := cached.SetupCache()

	// Setup Repository
	userRepo := mysql.SetupUserRepository(db)
	addressRepo := mysql.SetupAddressRepoRepository(db)
	productRepo := mysql.SetupProductRepoRepository(db)
	productTypeRepo := mysql.SetupProductTypeRepository(db)
	productImgRepo := mysql.SetupProductImgRepository(db)
	courierInfoRepo := mysql.SetupCourierInfoRepository(db)
	transactionRepo := mysql.SetupTransactionRepository(db)
	courierRepo := mysql.SetupCourierRepository(db)
	paymentInfo := mysql.SetupPaymentInfoRepository(db)

	// Setup Wrapper
	rajaongkirWrapper := rajaongkir.NewWrapper()
	mailjetWrapper := mailjet.NewWrapper()
	xenditWrapper := xendit.NewWrapper()
	dropboxWrapper := filebox.NewWrapper()

	// Setup Service
	userService := user.SetupUserService().
		SetUserRepo(userRepo).
		SetRajaOngkirWrapper(rajaongkirWrapper).
		SetRedisClient(cache).
		SetAddressRepo(addressRepo).
		SetMailjetWrapper(mailjetWrapper).
		Validate()

	productService := product.SetupProductService().
		SetProductRepo(productRepo).
		SetProductTypeRepo(productTypeRepo).
		SetProductImgRepo(productImgRepo).
		SetDropBoxWrapper(dropboxWrapper).
		SetCache(cache).
		Validate()

	transactionService := transaction.SetupTransactionService().
		SetProductRepo(productRepo).
		SetRajaongkirWrapper(rajaongkirWrapper).
		SetUserRepo(userRepo).
		SetXenditWrapper(xenditWrapper).
		SetCourierInfoRepo(courierInfoRepo).
		SetTransactionRepo(transactionRepo).
		SetCourierRepo(courierRepo).
		SetPaymentInfoRepo(paymentInfo).
		SetMailjetWrapper(mailjetWrapper).
		SetCache(cache).
		SetAddressRepo(addressRepo).
		Validate()

	// Setup Handler
	userHandler := handler.NewUserHandler(userService)
	healthCheckHandler := handler.NewHealthCheckHandler()
	productTypeHandler := handler.NewProductTypeHandler(productService)
	productHandler := handler.NewProductHandler(productService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Setup Router
	r := &router.Router{
		HealthCheck:        healthCheckHandler,
		UserHandler:        userHandler,
		ProductTypeHandler: productTypeHandler,
		ProductHandler:     productHandler,
		TransactionHandler: transactionHandler,
	}

	e := echo.New()
	r.NewRouter(e).Validate()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}

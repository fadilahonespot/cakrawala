package router

import (
	"github.com/fadilahonespot/cakrawala/server/handler"
	"github.com/fadilahonespot/cakrawala/server/middleware"
	"github.com/labstack/echo/v4"
)

type Router struct {
	HealthCheck        *handler.HealthCheckHandler
	UserHandler        *handler.UserHandler
	ProductTypeHandler *handler.ProductTypeHandler
	ProductHandler     *handler.ProductHandler
	TransactionHandler *handler.TransactionHandler
}

func (router *Router) Validate() {
	if router.HealthCheck == nil {
		panic("healty check handler is nil")
	}
	if router.UserHandler == nil {
		panic("user handler is nil")
	}
	if router.ProductTypeHandler == nil {
		panic("product type handler is nil")
	}
	if router.ProductHandler == nil {
		panic("product handler is nil")
	}
	if router.TransactionHandler == nil {
		panic("transaction handler is nil")
	}
}

func (router *Router) NewRouter(r *echo.Echo) *Router {
	middleware.SetupMiddleware(r)
	r.GET("/ping", router.HealthCheck.ServeHTTP)

	r.POST("/external/user/login", router.UserHandler.Login)
	r.POST("/external/user/register", router.UserHandler.Register)
	r.GET("/external/user/province", router.UserHandler.GetProvince)
	r.GET("/external/user/province/:provinceId/city", router.UserHandler.GetCity)
	r.GET("/external/user/verification", router.UserHandler.VerificationEmail)
	r.POST("/external/transaction/callback", router.TransactionHandler.CallbackTransaction)
	r.GET("/external/product/list", router.ProductHandler.GetAllProduct)
	r.GET("/external/product/:productId", router.ProductHandler.GetProductById)
	r.GET("/external/product/category", router.ProductTypeHandler.GetAllProductType)

	// Autorization
	au := r.Group("")
	au.Use(middleware.JwtMiddleware())
	au.POST("/user/resend-email", router.UserHandler.ResendEmailVerification)
	au.GET("/user/profile", router.UserHandler.GetProfile)
	au.PUT("/user/profile", router.UserHandler.UpdateProfile)

	au.GET("/transaction/bank", router.TransactionHandler.CheckAvailableBank)
	au.GET("/transaction/courier", router.TransactionHandler.CheckAvailabeCourier)
	au.POST("/transaction/shipping", router.TransactionHandler.CheckShipping)
	au.GET("/transaction/basket", router.TransactionHandler.GetBasket)
	au.POST("/transaction/basket", router.TransactionHandler.AddBasket)
	au.DELETE("/transaction/basket/:productId", router.TransactionHandler.DeleteBasket)
	au.POST("/transaction/checkout", router.TransactionHandler.CheckoutTransaction)
	au.GET("/transaction/history", router.TransactionHandler.GetHistory)
	au.GET("/transaction/detail/:transactionId", router.TransactionHandler.GetDetailTransaction)

	// Admin Onliy
	au.POST("/product/category", router.ProductTypeHandler.CreateProductType)
	au.PUT("/product/category/:productTypeId", router.ProductTypeHandler.UpdateProductType)
	au.POST("/product", router.ProductHandler.AddProduct)
	au.PUT("/product/:productId", router.ProductHandler.UpdateProduct)
	au.DELETE("/product/:productId", router.ProductHandler.DeleteProduct)
	au.POST("/product/images", router.ProductHandler.UploadProductImages)

	return router
}

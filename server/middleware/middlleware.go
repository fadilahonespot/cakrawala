package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fadilahonespot/cakrawala/utils/constrans"
	customerr "github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/fadilahonespot/cakrawala/utils/response"
	"github.com/fadilahonespot/library/logres"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cast"
)

func SetupMiddleware(server *echo.Echo) {
	server.Use(setLoggerMiddleware())
	server.Use(loggerMiddleware())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "token", echo.HeaderContentType, "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	server.HTTPErrorHandler = errorHandler
	server.Validator = &DataValidator{ValidatorData: validator.New()}
}

func errorHandler(err error, c echo.Context) {
	if c.Get("error-handled") != nil {
		return
	}

	c.Set("error-handled", true)

	code := http.StatusInternalServerError
	resp := response.ResponseError(code, "general error")

	if he, ok := err.(*customerr.ApplicationError); ok {
		resp.Code = he.ErrorCode
		resp.Message = he.Error()
	}

	request := c.Request()
	ctx := logres.SetErrorMessage(c.Request().Context(), err.Error())
	c.SetRequest(request.WithContext(ctx))

	c.JSON(resp.Code, resp)
}

func JwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, customerr.New(http.StatusUnauthorized, "signature not valid")
				}
				return []byte(constrans.JwtSecret), nil
			})

			if err != nil || !token.Valid {
				return customerr.New(http.StatusUnauthorized, "invalid token")
			}

			claims := token.Claims.(jwt.MapClaims)
			c.Set("user", claims)

			return next(c)
		}
	}
}

func loggerMiddleware() echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.TDR(c.Request().Context(), reqBody, resBody)
	})
}

func setLoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ctxLogger := logres.Context{
				ServiceName:    "cakrawala-app",
				ServiceVersion: "1.0.0",
				ServicePort:    cast.ToInt(os.Getenv("APP_PORT")),
				ThreadID:       uuid.New().String(),
				ReqMethod:      c.Request().Method,
				ReqURI:         c.Request().URL.String(),
				Header:         c.Request().Header,
			}

			request := c.Request()
			ctx := logres.SetCtxLogger(context.Background(), ctxLogger)
			c.SetRequest(request.WithContext(ctx))

			logger.Info(ctx, "Incoming Request")

			return next(c)
		}
	}
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	err := cv.ValidatorData.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var message string
		for _, err := range err.(validator.ValidationErrors) {
			message = fmt.Sprintf("field validator for input %v failed on the %v tag", err.Field(), err.ActualTag())
		}
		return errors.New(message)
	}
	return err
}

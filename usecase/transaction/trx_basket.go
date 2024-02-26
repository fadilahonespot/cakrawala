package transaction

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fadilahonespot/cakrawala/domain/entity"
	"github.com/fadilahonespot/cakrawala/usecase/transaction/model"
	"github.com/fadilahonespot/cakrawala/utils/errors"
	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func (s *defaultTransaction) GetAllProductBasket(ctx context.Context, userId string) (resp []model.BasketResponse, err error) {
	basketTrx, _ := s.transactionRepo.FindTransactionBasketByUserIdStatusPending(ctx, userId)
	if basketTrx == nil {
		logger.Info(ctx, "basket transaction is not data")
		return
	}

	dataProduct, err := s.transactionRepo.FindTransactionDetailByBasketId(ctx, strconv.Itoa(basketTrx.ID))
	if err != nil {
		logger.Error(ctx, "error find transaction detail basket", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(dataProduct); i++ {
		data := model.BasketResponse{
			ProductID:   dataProduct[i].ProductID,
			Name:        dataProduct[i].Product.Name,
			Weight:      dataProduct[i].Product.Weight,
			Price:       dataProduct[i].Product.Price,
			Qty:         dataProduct[i].Qty,
			TotalPrice:  dataProduct[i].Price,
			TotalWeight: dataProduct[i].Weight,
		}
		if len(dataProduct[i].Product.ProductImg) != 0 {
			data.Image = dataProduct[i].Product.ProductImg[0].Image
		}
		resp = append(resp, data)
	}

	return
}

func (s *defaultTransaction) AddProductBasket(ctx context.Context, userId string, req model.AddBasketRequest) (err error) {
	if req.Quantity <= 0 {
		logger.Error(ctx, "quantity must be greater than", err.Error())
		err = errors.New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	dataProduct, err := s.productRepo.FindByID(ctx, req.ProductID)
	if err != nil {
		logger.Error(ctx, "product not found", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if dataProduct.Stock <= 0 || !dataProduct.IsAvailable {
		logger.Error(ctx, "product not available")
		err = errors.New(http.StatusNotFound, "product not available")
		return
	}

	if dataProduct.Stock < req.Quantity {
		logger.Error(ctx, "quantity must be greater than product stock")
		err = errors.New(http.StatusBadRequest, "quantity must be greater than product stock")
		return
	}

	basketTrx, _ := s.transactionRepo.FindTransactionBasketByUserIdStatusPending(ctx, userId)
	if basketTrx.ID == 0 {
		reqBasket := entity.TransactionBasket{
			UserID:       cast.ToInt(userId),
			BasketStatus: entity.BasketPending,
		}
		err = s.transactionRepo.CreateTransactionBasket(ctx, &reqBasket)
		if err != nil {
			logger.Error(ctx, "error creating transaction basket", err.Error())
			err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		basketTrx.ID = reqBasket.ID
	}
	trxProduct, err := s.transactionRepo.FindTransactionDetailByProductIdAndBasketId(ctx, strconv.Itoa(req.ProductID), strconv.Itoa(basketTrx.ID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tx := s.productRepo.BeginTrans(ctx)

			trxDetail := entity.TransactionDetail{
				TransactionBasketID: basketTrx.ID,
				ProductID:           req.ProductID,
				Qty:                 req.Quantity,
				Price:               dataProduct.Price * float64(req.Quantity),
				Weight:              dataProduct.Weight * req.Quantity,
			}
			err = s.transactionRepo.CreateTransactionDetail(ctx, tx, &trxDetail)
			if err != nil {
				tx.Rollback()
				logger.Error(ctx, "failed create transaction detail", err.Error())
				err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}

			tx.Commit()
			return
		}
		logger.Error(ctx, "error find transaction detail by product id", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	tx := s.productRepo.BeginTrans(ctx)
	trxDetail := entity.TransactionDetail{
		ID:                  trxProduct.ID,
		TransactionBasketID: basketTrx.ID,
		ProductID:           req.ProductID,
		Qty:                 trxProduct.Qty + req.Quantity,
		Price:               trxProduct.Price + (dataProduct.Price * float64(req.Quantity)),
		Weight:              trxProduct.Weight + (dataProduct.Weight * req.Quantity),
		CreatedAt:           trxProduct.CreatedAt,
	}
	err = s.transactionRepo.UpdateTransactionDetail(ctx, tx, &trxDetail)
	if err != nil {
		tx.Rollback()
		logger.Error(ctx, "error update transaction detail", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	tx.Commit()
	return
}

func (s *defaultTransaction) DeleteProductBasket(ctx context.Context, userId, productId string) (err error) {
	basketTrx, err := s.transactionRepo.FindTransactionBasketByUserIdStatusPending(ctx, userId)
	if err != nil {
		logger.Error(ctx, "error find transaction basket by user id", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	trxDetail, err := s.transactionRepo.FindTransactionDetailByProductIdAndBasketId(ctx, productId, strconv.Itoa(basketTrx.ID))
	if err != nil {
		logger.Error(ctx, "error find transaction detail", err.Error())
		err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	tx := s.productRepo.BeginTrans(ctx)
	err = s.transactionRepo.DeleteTransactionDetail(ctx, tx, trxDetail.ID)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			logger.Error(ctx, "product id not found", err.Error())
			err = errors.New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		logger.Error(ctx, "error delete transaction detail", err.Error())
		err = errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	tx.Commit()
	return
}

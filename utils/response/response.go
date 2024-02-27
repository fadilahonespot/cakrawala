package response

import (
	"math"
	"net/http"

	"github.com/fadilahonespot/cakrawala/utils/paginate"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responsePagination struct {
	Response
	Pagination paginate.ItemPages `json:"pagination"`
}

func ResponseSuccess(data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}

func ResponseSuccessWithMessage(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    struct{}{},
	}
}


func ResponseProcees(data interface{}) Response {
	return Response{
		Code:    http.StatusProcessing,
		Message: "Request has been successfully processed",
		Data:    data,
	}
}

func ResponseError(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    struct{}{},
	}
}

func HandleSuccessWithPagination(totalItems float64, params paginate.Pagination, data interface{}) responsePagination {
	var totalPage float64 = 1
	if params.Limit != 0 && params.Page != 0 {
		res := totalItems / float64(params.Limit)
		totalPage = math.Ceil(res)
	}

	resp := responsePagination{
		Response: Response{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    data,
		},
		Pagination: paginate.ItemPages{
			TotalData: int64(totalItems),
			TotalPage: int64(totalPage),
			Page:      int64(params.Page),
			Limit:     int64(params.Limit),
		},
	}
	return resp
}

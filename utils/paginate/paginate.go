package paginate

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Pagination struct {
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
	Key      string `query:"key"`
	Value    string `query:"value"`
	In       string `query:"in"`
	Order    string `query:"order"`
	Sort     string `query:"sort"`
}

type ItemPages struct {
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	TotalData int64 `json:"total_data"`
	TotalPage int64 `json:"total_page"`
}

func Paginate(page, length int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case length > 50:
			length = 50
		case length <= 0:
			length = 10
		}

		offset := (page - 1) * length
		return db.Offset(offset).Limit(length)
	}
}

func GetParams(c echo.Context) (Pagination, error) {
	params := Pagination{
		Page:     cast.ToInt(c.QueryParam("page")),
		Limit:    cast.ToInt(c.QueryParam("limit")),
		Key:      c.QueryParam("key"),
		Value:    c.QueryParam("value"),
		In:       c.QueryParam("in"),
		Order:    c.QueryParam("order"),
		Sort:     c.QueryParam("sort"),
	}

	err := c.Validate(params)
	if err != nil {
		return Pagination{}, errors.New("error validating pagination")
	}
	return params, nil
}
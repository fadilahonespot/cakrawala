package utils

type Pagination struct {
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
	Key      string `query:"key"`
	Value    string `query:"value"`
	In       string `query:"in"`
	FromDate string `query:"from" validate:"number"`
	ToDate   string `query:"to" validate:"number"`
	Order    string `query:"order"`
	Sort     string `query:"sort"`
}
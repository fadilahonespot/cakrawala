package model

type Product struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Code            string   `json:"code"`
	IsAvailable     bool     `json:"isAvailable"`
	Description     string   `json:"description"`
	ProductTypeID   int      `json:"categoryId"`
	ProductTypeName string   `json:"category"`
	Weight          int      `json:"weight"`
	Stock           int      `json:"stock"`
	Price           float64  `json:"price"`
	Sold            int      `json:"sold"`
	Images          []string `json:"images"`
}

type GetAllProductTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetAllProductResponse struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	IsAvailable     bool    `json:"isAvailable"`
	ProductTypeID   int     `json:"categoryId"`
	ProductTypeName string  `json:"category"`
	Price           float64 `json:"price"`
	Image           string  `json:"image"`
	Sold            int     `json:"sold"`
}

type UploadImageResponse struct {
	Images []string `json:"images"`
}

type GenerateTextResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
type GenerateRequest struct {
	ProductId int `json:"productId" validate:"required"`
}

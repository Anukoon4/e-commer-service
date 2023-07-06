package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductRequest struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

type ProductData struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	ProductData
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductModel struct {
	gorm.Model
	Name  string `gorm:"column:name" json:"name"`
	Price int    `gorm:"column:price" json:"price"`
}

func (ctl *ProductModel) ToModels(req *ProductData) {
	if v := req.Name; v != "" {
		ctl.Name = v
	}
	if v := req.Price; v != 0 {
		ctl.Price = v
	}

}

func (ctl *ProductModel) ToResponse() *ProductResponse {
	response := ProductResponse{}
	if v := ctl.ID; v != 0 {
		response.Id = int(v)
	}
	if v := ctl.Name; v != "" {
		response.Name = v
	}
	if v := ctl.Price; v != 0 {
		response.Price = v
	}
	response.CreatedAt = ctl.CreatedAt
	response.UpdatedAt = ctl.UpdatedAt
	return &response
}

func (ctl *ProductModel) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&ProductModel{})
}

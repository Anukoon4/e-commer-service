package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderProductData struct {
	ProductData
	Count int `json:"count"`
}

type OrderProductResponse struct {
	Id        int       `json:"id,omitempty"`
	ProductId int       `json:"product_id"`
	OrderId   int       `json:"order_id"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderProductModel struct {
	gorm.Model
	ProductId int           `gorm:"column:product_id" json:"product_id"`
	Product   *ProductModel `gorm:"foreignKey:ProductId"`
	OrderId   int           `gorm:"column:order_id" json:"order_id"`
	Count     int           `gorm:"column:count" json:"count"`
}

func (ctl *OrderProductModel) ToModels(req *OrderProductData) {
	if v := req.Id; v != 0 {
		ctl.ProductId = v
	}
	if v := req.Count; v != 0 {
		ctl.Count = v
	}
}

func (ctl *OrderProductModel) ToResponse() *OrderProductData {
	res := OrderProductData{}
	if v := ctl.ID; v != 0 {
		res.Id = int(v)
	}
	if v := ctl.Count; v != 0 {
		res.Count = v
	}
	if v := ctl.Product; v != nil {
		res.ProductData = v.ToResponse().ProductData
	}
	return &res
}

func (ctl *OrderProductModel) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&OrderProductModel{})
}

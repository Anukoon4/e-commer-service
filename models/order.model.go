package models

import (
	"gorm.io/gorm"
)

type OrderRequest struct {
	Id       int                 `json:"id,omitempty"`
	UserId   int                 `json:"user_id"`
	Products []*OrderProductData `json:"products"`
}

type OrderResponse struct {
	Id       int                 `json:"id,omitempty"`
	UserId   int                 `json:"user_id"`
	Products []*OrderProductData `json:"products"`
}

type OrderModel struct {
	gorm.Model
	UserId   int
	User     UserModel            `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Products []*OrderProductModel `gorm:"foreignKey:OrderId"`
}

func (ctl *OrderModel) ToModels(req *OrderRequest) {

	if v := req.UserId; v != 0 {
		ctl.UserId = v
	}

	if v := req.Products; len(v) != 0 {
		products := []*OrderProductModel{}
		for _, item := range v {
			product := OrderProductModel{}
			product.ToModels(item)
			products = append(products, &product)
		}
		ctl.Products = products
	}
}

func (ctl *OrderModel) ToResponse() *OrderResponse {
	res := OrderResponse{}
	if v := ctl.ID; v != 0 {
		res.Id = int(v)
	}
	if v := ctl.UserId; v != 0 {
		res.UserId = v
	}

	if v := ctl.Products; len(v) != 0 {
		items := []*OrderProductData{}
		for _, item := range v {
			items = append(items, item.ToResponse())
		}

		res.Products = items
	}
	return &res
}

func (ctl *OrderModel) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&OrderModel{})
}

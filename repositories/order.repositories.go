package repositories

import (
	"context"

	config "github.com/anukoon4/e-commerce-service/configs"
	"github.com/anukoon4/e-commerce-service/models"
	"gorm.io/gorm"
)

type OrderRepositories struct {
	mysql config.MySQL
}

func (r *OrderRepositories) QueryBinding(req *models.OrderRequest) (*gorm.DB, error) {
	qr := r.mysql.DB().Model(&models.OrderModel{})
	if v := req.Id; v != 0 {
		qr = qr.Where("id = ?", v)
	}
	qr.Preload("Products")
	qr.Preload("Products.Product")

	return qr, nil
}

func (r *OrderRepositories) Create(ct context.Context, req *models.OrderRequest) (*models.OrderModel, error) {
	productModel := models.OrderModel{}
	productModel.ToModels(req)
	qr := r.mysql.DB().Model(&models.OrderModel{})
	if err := qr.Create(&productModel).Error; err != nil {
		return nil, err
	}
	return &productModel, nil
}

func (r *OrderRepositories) FindOne(ct context.Context, req *models.OrderRequest) (*models.OrderModel, error) {
	productModel := models.OrderModel{}
	qr, err := r.QueryBinding(req)
	if err != nil {
		return nil, err
	}
	if err := qr.First(&productModel).Error; err != nil {
		return nil, err
	}
	return &productModel, nil
}

func (r *OrderRepositories) FindOrderByUserId(ct context.Context, req *models.OrderRequest) ([]*models.OrderModel, error) {
	order := []*models.OrderModel{}
	qr, err := r.QueryBinding(req)
	if err != nil {
		return nil, err
	}
	if err := qr.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepositories) Delete(ct context.Context, req *models.OrderRequest) (*models.OrderModel, error) {
	if model, err := r.FindOne(ct, req); err != nil {
		return nil, err
	} else {
		model.ToModels(req)
		qr := r.mysql.DB().Model(&models.OrderModel{})
		qr = qr.Where("id = ? AND user_id = ?", req.Id, req.UserId)
		if err := qr.Delete(&model).Error; err != nil {
			return nil, err
		}
		return model, nil
	}
}

package repositories

import (
	"context"

	config "github.com/anukoon4/e-commerce-service/configs"
	"github.com/anukoon4/e-commerce-service/models"
	"github.com/anukoon4/e-commerce-service/utils"
	"gorm.io/gorm"
)

type ProductRepositories struct {
	mysql config.MySQL
}

func (r *ProductRepositories) QueryBinding(req *models.ProductRequest) (*gorm.DB, error) {
	qr := r.mysql.DB().Model(&models.ProductModel{})
	if v := req.Id; v != 0 {
		qr = qr.Where("id = ?", v)
	}
	if v := req.Name; v != "" {
		qr = qr.Where("name = ?", v)
	}
	if v := req.Price; v != 0 {
		qr = qr.Where("price = ?", v)
	}
	return qr, nil
}

func (r *ProductRepositories) Create(ct context.Context, req *models.ProductData) (*models.ProductModel, error) {
	productModel := models.ProductModel{}
	productModel.ToModels(req)
	qr := r.mysql.DB().Model(&models.ProductModel{})
	if err := qr.Create(&productModel).Error; err != nil {
		return nil, err
	}
	return &productModel, nil
}

func (r *ProductRepositories) FindOne(ct context.Context, req *models.ProductRequest) (*models.ProductModel, error) {
	productModel := models.ProductModel{}
	qr, err := r.QueryBinding(req)
	if err != nil {
		return nil, err
	}
	if err := qr.First(&productModel).Error; err != nil {
		return nil, err
	}
	return &productModel, nil
}

func (r *ProductRepositories) FindAll(ct context.Context, req *models.ProductRequest) ([]*models.ProductModel, *utils.Paginator, error) {
	productsModel := []*models.ProductModel{}
	qr, err := r.QueryBinding(req)
	if err != nil {
		return nil, nil, err
	}
	meta, err := utils.Paging(&utils.PagingConfig{
		DB:      qr,
		Page:    req.Page,
		PerPage: req.PerPage,
	}, &productsModel)

	return productsModel, meta, err
}

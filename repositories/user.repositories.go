package repositories

import (
	"context"

	config "github.com/anukoon4/e-commerce-service/configs"
	"github.com/anukoon4/e-commerce-service/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositories struct {
	mysql config.MySQL
}

func (r *UserRepositories) QueryBinding(req *models.UserRequest) (*gorm.DB, error) {
	qr := r.mysql.DB().Model(new(models.UserModel))
	if v := req.Id; v != 0 {
		qr = qr.Where("id = ?", v)
	}
	if v := req.Name; v != "" {
		qr = qr.Where("name = ?", v)
	}
	if v := req.Email; v != "" {
		qr = qr.Where("email = ?", v)
	}
	if v := req.Password; v != "" {
		qr = qr.Where("password = ?", v)
	}
	if v := req.Phone; v != "" {
		qr = qr.Where("phone = ?", v)
	}
	return qr, nil
}

func (r *UserRepositories) Create(ct context.Context, req *models.UserRequest) (*models.UserModel, error) {
	useModel := models.UserModel{}
	useModel.ToModels(req)
	pwdEncrypted, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	useModel.Password = string(pwdEncrypted)
	qr := r.mysql.DB().Model(&models.UserModel{})
	if err := qr.Create(&useModel).Error; err != nil {
		return nil, err
	}
	return &useModel, nil
}

func (r *UserRepositories) FindOne(ct context.Context, req *models.UserRequest) (*models.UserModel, error) {
	userModel := models.UserModel{}
	qr, err := r.QueryBinding(req)
	if err != nil {
		return nil, err
	}
	if err := qr.First(&userModel).Error; err != nil {
		return nil, err
	}
	return &userModel, nil
}

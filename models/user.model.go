package models

import (
	"gorm.io/gorm"
)

type UserRequest struct {
	Id       int    `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

type UserModel struct {
	gorm.Model
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
	Phone    string `gorm:"column:phone" json:"phone"`
}

func (ctl *UserModel) ToModels(req *UserRequest) {
	if v := req.Email; v != "" {
		ctl.Email = v
	}
	if v := req.Password; v != "" {
		ctl.Password = v
	}
	if v := req.Name; v != "" {
		ctl.Name = v
	}
	if v := req.Phone; v != "" {
		ctl.Phone = v
	}
}

func (ctl *UserModel) ToResponse() *UserRequest {
	response := UserRequest{}
	if v := ctl.ID; v != 0 {
		response.Id = int(v)
	}
	if v := ctl.Email; v != "" {
		response.Email = v
	}
	if v := ctl.Name; v != "" {
		response.Name = v
	}
	if v := ctl.Phone; v != "" {
		response.Phone = v
	}
	return &response
}

func (ctl *UserModel) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
}

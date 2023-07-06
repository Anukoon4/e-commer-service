package utils

import (
	"math"

	"gorm.io/gorm"
)

type PagingConfig struct {
	DB      *gorm.DB
	Page    int
	PerPage int
	ShowSQL bool
	All     bool
}

// Paginator paging response
type Paginator struct {
	TotalRecord int `json:"total_record"`
	TotalPage   int `json:"total_page"`
	PerPage     int `json:"PerPage"`
	Page        int `json:"page"`
	Offset      int `json:"offset"`
}

func Paging(p *PagingConfig, result interface{}) (*Paginator, error) {
	db := p.DB
	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}

	var paginator Paginator
	var count int64
	var offset int

	model := db.Statement.Model
	if model == nil {
		model = result
	}

	if err := db.Model(model).Count(&count).Error; err != nil {
		return nil, err
	}

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.PerPage
	}

	if err := db.Limit(p.PerPage).Offset(offset).Find(result).Error; err != nil {
		return nil, err
	}

	paginator.TotalRecord = int(count)
	paginator.Page = p.Page
	paginator.Offset = offset
	paginator.PerPage = p.PerPage
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.PerPage)))
	return &paginator, nil
}

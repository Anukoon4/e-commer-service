package handlers

import (
	"net/http"

	"github.com/anukoon4/e-commerce-service/constants"
	"github.com/anukoon4/e-commerce-service/models"
	"github.com/anukoon4/e-commerce-service/repositories"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	productRepo repositories.ProductRepositories
}

func (ctl *ProductHandlers) Finds(c *gin.Context) {
	var productReq models.ProductRequest

	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userRes, meta, err := ctl.productRepo.FindAll(c, &productReq)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": constants.NOT_FOUND})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}

	userResList := []*models.ProductResponse{}
	for _, item := range userRes {
		userResList = append(userResList, item.ToResponse())
	}

	c.JSON(http.StatusOK, gin.H{"data": userResList, "meta": meta})
}

func (ctl *ProductHandlers) Find(c *gin.Context) {
	var productReq models.ProductRequest

	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userRes, err := ctl.productRepo.FindOne(c, &productReq)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": constants.NOT_FOUND})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userRes.ToResponse()})
}

func (ctl *ProductHandlers) Create(c *gin.Context) {
	var productData models.ProductData

	if err := c.ShouldBindJSON(&productData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userRes, err := ctl.productRepo.Create(c, &productData)
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userRes.ToResponse()})
}

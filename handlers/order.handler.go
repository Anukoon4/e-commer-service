package handlers

import (
	"net/http"

	"github.com/anukoon4/e-commerce-service/constants"
	"github.com/anukoon4/e-commerce-service/models"
	"github.com/anukoon4/e-commerce-service/repositories"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type OrderHandlers struct {
	orderRepo repositories.OrderRepositories
}

func (ctl *OrderHandlers) Create(c *gin.Context) {
	var orderReq models.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userId := int(c.MustGet("user_id").(float64))
	orderReq.UserId = userId
	res, err := ctl.orderRepo.Create(c, &orderReq)
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (ctl *OrderHandlers) Find(c *gin.Context) {
	var orderReq models.OrderRequest

	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	orderRes, err := ctl.orderRepo.FindOne(c, &orderReq)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": constants.NOT_FOUND})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderRes.ToResponse()})
}

func (ctl *OrderHandlers) MyOrder(c *gin.Context) {
	userId := int(c.MustGet("user_id").(float64))

	orderRequest := models.OrderRequest{
		UserId: userId,
	}
	orderRes, err := ctl.orderRepo.FindOrderByUserId(c, &orderRequest)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": constants.NOT_FOUND})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}
	_orderRer := []*models.OrderResponse{}
	for _, item := range orderRes {
		_orderRer = append(_orderRer, item.ToResponse())
	}
	c.JSON(http.StatusOK, gin.H{"data": _orderRer})
}

func (ctl *OrderHandlers) Delete(c *gin.Context) {
	var orderReq models.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userId := int(c.MustGet("user_id").(float64))
	orderReq.UserId = userId

	res, err := ctl.orderRepo.Delete(c, &orderReq)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": constants.NOT_FOUND})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

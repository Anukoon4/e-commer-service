package handlers

import (
	"net/http"

	"github.com/anukoon4/e-commerce-service/constants"
	"github.com/anukoon4/e-commerce-service/models"
	"github.com/anukoon4/e-commerce-service/repositories"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	userRepo repositories.UserRepositories
}

func (ctl *UserHandlers) Profile(c *gin.Context) {
	userId := int(c.MustGet("user_id").(float64))
	req := models.UserRequest{
		Id: userId,
	}
	userRes, err := ctl.userRepo.FindOne(c, &req)
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

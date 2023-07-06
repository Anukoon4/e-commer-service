package handlers

import (
	"net/http"
	"strings"

	"github.com/anukoon4/e-commerce-service/constants"
	"github.com/anukoon4/e-commerce-service/models"
	"github.com/anukoon4/e-commerce-service/repositories"
	service "github.com/anukoon4/e-commerce-service/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	userRepo repositories.UserRepositories
}

func (ctl *AuthHandlers) Register(c *gin.Context) {
	var userRequest models.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//check duplicate
	req := models.UserRequest{
		Email: userRequest.Email,
	}
	userRes, _ := ctl.userRepo.FindOne(c, &req)
	if userRes != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email duplicate"})
		return
	}

	userRes, err := ctl.userRepo.Create(c, &userRequest)
	if err != nil {
		switch err.Error() {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userRes})
}

func (ctl *AuthHandlers) Login(c *gin.Context) {
	var userRequest models.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	email := strings.Trim(userRequest.Email, "")
	req := models.UserRequest{
		Email: email,
	}
	account, err := ctl.userRepo.FindOne(c, &req)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.INTERNAL_SERVER_ERROR})
		}
		return
	}

	if account != nil {
		err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(userRequest.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		tokenResult := service.GenToken(account)
		c.JSON(http.StatusOK, gin.H{"token": tokenResult})
	}

}

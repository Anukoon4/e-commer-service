package routes

import (
	"github.com/anukoon4/e-commerce-service/handlers"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authHandlers handlers.AuthHandlers
}

func (c *AuthRoutes) RegisterAuthRouting(router *gin.RouterGroup) {
	router.POST("/register", c.authHandlers.Register)
	router.POST("/login", c.authHandlers.Login)
}

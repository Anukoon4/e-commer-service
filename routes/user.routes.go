package routes

import (
	"github.com/anukoon4/e-commerce-service/handlers"
	"github.com/anukoon4/e-commerce-service/middleware"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandlers  handlers.UserHandlers
	orderHandlers handlers.OrderHandlers
}

func (c *UserRoutes) RegisterUserRouting(router *gin.RouterGroup) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/profile", c.userHandlers.Profile)
	router.GET("/my-order", c.orderHandlers.MyOrder)
}

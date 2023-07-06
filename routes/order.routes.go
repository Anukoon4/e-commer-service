package routes

import (
	"github.com/anukoon4/e-commerce-service/handlers"
	"github.com/anukoon4/e-commerce-service/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRoutes struct {
	orderHandlers handlers.OrderHandlers
}

func (c *OrderRoutes) RegisterOrderRouting(router *gin.RouterGroup) {
	router.Use(middleware.AuthMiddleware())
	router.POST("/create", c.orderHandlers.Create)
	router.POST("/find", c.orderHandlers.Find)
	router.DELETE("/delete", c.orderHandlers.Delete)
}

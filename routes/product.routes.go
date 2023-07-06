package routes

import (
	"github.com/anukoon4/e-commerce-service/handlers"
	"github.com/anukoon4/e-commerce-service/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	productHandlers handlers.ProductHandlers
}

func (c *ProductRoutes) RegisterProductRouting(router *gin.RouterGroup) {
	router.Use(middleware.AuthMiddleware())
	router.POST("/create", c.productHandlers.Create)
	router.POST("/finds", c.productHandlers.Finds)
	router.POST("/find", c.productHandlers.Find)
}

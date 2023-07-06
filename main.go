package main

import (
	"fmt"
	"os"

	config "github.com/anukoon4/e-commerce-service/configs"
	"github.com/anukoon4/e-commerce-service/models"
	routes "github.com/anukoon4/e-commerce-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	user := models.UserModel{}
	product := models.ProductModel{}
	order := models.OrderModel{}
	orderProduct := models.OrderProductModel{}
	user.AutoMigrate(db)
	product.AutoMigrate(db)
	order.AutoMigrate(db)
	orderProduct.AutoMigrate(db)
}

func main() {
	godotenv.Load()
	db := config.Init()
	Migrate(db)
	router := gin.Default()
	router.Use(core())
	port := os.Getenv("APP_PORT")
	portStr := fmt.Sprintf(":%s", port)

	authRoute := routes.AuthRoutes{}
	userRoute := routes.UserRoutes{}
	productRoute := routes.ProductRoutes{}
	orderRoute := routes.OrderRoutes{}
	api := router.Group("/")
	authRoute.RegisterAuthRouting(api.Group("/auth"))
	userRoute.RegisterUserRouting(api.Group("/user"))
	productRoute.RegisterProductRouting(api.Group("/product"))
	orderRoute.RegisterOrderRouting(api.Group("/order"))

	err := router.Run(portStr)
	if err != nil {
		panic("[Error] failed to start Gin server: " + err.Error())
	}
}

func core() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
		c.Next()
	}
}

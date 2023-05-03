package main

import (
	"ecommerce/controller"
	"ecommerce/database"
	"ecommerce/middleware"
	"ecommerce/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "ecommerce/docs"

	swaggerfile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MarketPlace
// @version 1.0.0
// @description Api do marketPlace.

// @contact.name Lucas Euzebio

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Address{})
}

func serveApplication() {
	router := gin.Default()

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfile.Handler, url))

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/address", controller.AddAddress)
	protectedRoutes.GET("/address", controller.GetAllAddress)
	protectedRoutes.GET("/user", controller.Getuser)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

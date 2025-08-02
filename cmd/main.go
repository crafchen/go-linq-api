package main

import (
	"go-linq-api/internal/controllers"
	dbContext "go-linq-api/internal/db"
	"go-linq-api/internal/repositories"
	"go-linq-api/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := dbContext.ConnectMSSQL()
	if err != nil {
		log.Fatal("Cannot connect database:", err)
	}

	// Khởi tạo repository, service, controller
	provinceRepo := repositories.NewProvinceRepository(db)
	provinceService := services.NewProvinceService(provinceRepo)
	provinceController := controllers.NewProvinceController(provinceService)

	// Ward
	wardRepo := repositories.NewWardRepository(db)
	wardService := services.NewWardService(wardRepo)
	wardController := controllers.NewWardController(wardService)

	r := gin.Default()
	provinceController.RegisterRoutes(r)
	wardController.RegisterRoutes(r)

	log.Println("Server is running at :8080")
	r.Run(":8080")
}

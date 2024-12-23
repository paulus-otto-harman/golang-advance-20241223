package main

import (
	"20241223/model"
	"20241223/route"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=20241223b port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&model.User{})

	router := gin.Default()
	route.AuthRoutes(router, db)

	router.Run(":8081")
}

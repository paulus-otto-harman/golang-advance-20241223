package main

import (
	//"20241223/model"
	"20241223/route"
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

func main() {

	router := gin.Default()
	route.AuthRoutes(router)

	router.Run(":8081")
}

package route

import (
	"20241223/controller"
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

func AuthRoutes(router *gin.Engine) {
	authController := controller.NewAuthController()

	authRoutes := router.Group("/")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/authentication", authController.Authentication)
	}
}

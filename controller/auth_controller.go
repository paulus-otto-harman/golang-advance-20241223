package controller

import (
	"20241223/model"
	"log"

	//"20241223/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	//"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	//repo repository.AuthRepository
}

func NewAuthController(db *gorm.DB) *AuthController {
	//return &AuthController{repo: repository.NewAuthRepository(db)}
	return &AuthController{}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	//var login model.Login
	//if err := c.ShouldBindJSON(&login); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
	//	return
	//}
	//
	//if err := ctrl.repo.Login(&login); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "login failed"})
	//	return
	//}
	log.Println("login")
	log.Println("request from IP address", c.ClientIP())
	c.JSON(http.StatusOK, response{Success: true, Message: "login success"})
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response{Success: false, Message: err.Error()})
		return
	}
	log.Println("register")
	log.Println("request from IP address", c.ClientIP())
	c.JSON(http.StatusOK, response{Success: true, Message: "user registered", Data: struct {
		Name     string `json:"name"`
		Username string `json:"username"`
	}{
		Name:     user.Name,
		Username: user.Username,
	}})
}

func (ctrl *AuthController) Authentication(c *gin.Context) {
	log.Println("authenticate user")
	log.Println("request from IP address", c.ClientIP())
	c.JSON(http.StatusOK, response{Success: true, Message: "user authenticated"})
}

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

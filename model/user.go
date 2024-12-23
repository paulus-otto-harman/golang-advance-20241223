package model

type User struct {
	Name            string `json:"name" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=4,max=16"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

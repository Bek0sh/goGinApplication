package controller

import (
	"net/http"
	gormdb "project2/gormDB"
	"project2/models"
	"project2/password"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := password.HashPassword(input.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Name:     input.Name,
		Surname:  input.Surname,
		Email:    input.Email,
		Password: hashedPassword,
	}

	created_user, err := user.CreateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,
		gin.H{
			"created_user": created_user,
		})
}

func Login(c *gin.Context) {
	var inputUser models.LoginUserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	if err := gormdb.DB.First(&user, "email=?", inputUser.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Do not find user",
		})
		return
	}

	if err := password.ComparePassword(inputUser.Password, user.Password); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": "password does not match",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/spendwise/api/v1/dto"
	"github.com/kweku-xvi/spendwise/api/v1/models"
	"github.com/kweku-xvi/spendwise/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body dto.SignUpRequest
	
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		MiddleName: body.MiddleName,
		LastName: body.LastName,
		Email: body.Email,
		Username: body.Username,
		Password: string(passwordHash),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":result.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":user,
	})
}
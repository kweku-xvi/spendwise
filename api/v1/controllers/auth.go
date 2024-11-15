package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kweku-xvi/spendwise/api/v1/dto"
	"github.com/kweku-xvi/spendwise/api/v1/models"
	"github.com/kweku-xvi/spendwise/internal/config"
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

func SignIn(c *gin.Context) {
	var body dto.SignInRequest

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})
		return
	}

	var userFound models.User
	database.DB.Where("email=?", body.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"user not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid password",
		})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(config.ENV.JWTSecret))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":token,
	})
}

func GetUserProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")

	c.JSON(http.StatusOK, gin.H{
		"user":user,
	})
}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/spendwise/api/v1/controllers"
	"github.com/kweku-xvi/spendwise/internal/database"
)

func init() {
	database.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users/signup", controllers.SignUp)
	}

	r.Run()
}

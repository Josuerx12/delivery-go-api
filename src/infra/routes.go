package Routes

import (
	"delivery-go-api/src/infra/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, database *gorm.DB) {
	userController := controllers.NewUserController(database)

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/users", userController.CreateUser)
}
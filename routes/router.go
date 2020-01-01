package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tara/services"
)

func InitRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "hello, world!"})
	})

	userAPI := router.Group("/api/v1/users")
	
	{
		userAPI.GET("/", services.GetUsers)
		userAPI.GET("/:id", services.GetUser)
		userAPI.POST("/", services.CreateUser)
		userAPI.POST("/:id", services.UpdateUser)
		userAPI.DELETE("/:id", services.DeleteUser)
	}

	return router
}

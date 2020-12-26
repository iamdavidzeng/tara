package routes

import (
	"net/http"
	"tara/api/services"

	"github.com/gin-gonic/gin"
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

	postAPI := router.Group("/api/v1/posts")
	{
		postAPI.GET("/", services.GetPosts)
		postAPI.GET("/:id", services.GetPost)
		postAPI.POST("/", services.CreatePost)
		postAPI.POST("/:id", services.UpdatePost)
		postAPI.DELETE("/:id", services.DeletePost)
	}

	return router
}

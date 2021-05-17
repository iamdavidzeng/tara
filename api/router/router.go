package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdavidzeng/tara/api/services/posts"
	"github.com/iamdavidzeng/tara/api/services/users"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "hello, world!"})
	})

	userAPI := router.Group("/api/v1/users")
	{
		userAPI.GET("", users.List)
		userAPI.GET("/:id", users.Get)
		userAPI.POST("", users.New)
		userAPI.POST("/:id", users.Update)
		userAPI.DELETE("/:id", users.Delete)
	}

	postAPI := router.Group("/api/v1/posts")
	{
		postAPI.GET("", posts.List)
		postAPI.GET("/:id", posts.Get)
		postAPI.POST("", posts.New)
		postAPI.POST("/:id", posts.Update)
		postAPI.DELETE("/:id", posts.Delete)
	}

	return router
}

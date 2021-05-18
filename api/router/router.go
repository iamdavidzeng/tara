package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdavidzeng/tara/api/services"
	"github.com/iamdavidzeng/tara/api/services/posts"
	"github.com/iamdavidzeng/tara/api/services/users"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "hello, world!"})
	})

	userGroup := router.Group("/api/v1/users")
	registerRoute(users.UserOperator, userGroup)

	postGroup := router.Group("/api/v1/posts")
	registerRoute(posts.PostOperator, postGroup)

	return router
}

func registerRoute(m services.ModelOperator, r *gin.RouterGroup) {
	r.GET("", m.List)
	r.GET("/:id", m.Get)
	r.POST("", m.New)
	r.POST("/:id", m.Update)
	r.DELETE("/:id", m.Del)
}

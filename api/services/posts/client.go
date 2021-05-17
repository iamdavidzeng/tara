package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdavidzeng/tara/api/models"
	"github.com/iamdavidzeng/tara/api/schemas"
	"github.com/iamdavidzeng/tara/internal/db"
)

func New(c *gin.Context) {
	var data schemas.PostSchema
	c.BindJSON(&data)

	post := models.Posts{UserID: data.UserID, Title: data.Title, Content: data.Content, Published: data.Published}
	db.D.Storage.Create(&post)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Create successfully!"})
}

func Get(c *gin.Context) {
	var post models.Posts
	postID := c.Param("id")

	db.D.Storage.First(&post, postID)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Post Not Found!"})
		return
	}

	data := schemas.PostSchema{ID: post.ID, Title: post.Title, Content: post.Content}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func List(c *gin.Context) {
	var posts []models.Posts
	var data []schemas.PostSchema

	db.D.Storage.Find(&posts)
	if len(posts) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusUnauthorized, "data": "Post Not Found!"})
		return
	}

	for _, post := range posts {
		data = append(data, schemas.PostSchema{ID: post.ID, Title: post.Title, Content: post.Content})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func Update(c *gin.Context) {
	var post models.Posts
	postID := c.Param("id")

	db.D.Storage.First(&post, postID)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Post Not Found!"})
		return
	}

	var data schemas.PostSchema
	c.BindJSON(&data)

	db.D.Storage.Model(&post).Updates(data)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Update successfully!"})
}

func Delete(c *gin.Context) {
	var post models.Posts
	postID := c.Param("id")

	db.D.Storage.First(&post, postID)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Post Not Found!"})
		return
	}

	db.D.Storage.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Delete successfully!"})
}

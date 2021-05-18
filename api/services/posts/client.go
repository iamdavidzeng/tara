package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdavidzeng/tara/api/services"
	"github.com/iamdavidzeng/tara/internal/db"
)

type Posts struct {
	services.GormBase
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published,omitempty"`
}

var PostOperator *Posts = &Posts{}

func (p Posts) New(c *gin.Context) {
	c.BindJSON(&p)

	db.D.Storage.Create(&p)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Create successfully!"})
}

func (p Posts) Get(c *gin.Context) {
	id := c.Param("id")

	db.D.Storage.First(&p, id)

	if p.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusNotFound, "data": "Post Not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": p})
}

func (p Posts) List(c *gin.Context) {
	posts := []Posts{}

	db.D.Storage.Find(&posts)
	if len(posts) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": posts})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": posts})
}

func (p Posts) Update(c *gin.Context) {
	id := c.Param("id")

	db.D.Storage.First(&p, id)
	if p.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Post Not Found!"})
		return
	}

	c.BindJSON(&p)

	db.D.Storage.Model(&p).Updates(p)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Update successfully!"})
}

func (p Posts) Del(c *gin.Context) {
	id := c.Param("id")

	db.D.Storage.First(&p, id)
	if p.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Post Not Found!"})
		return
	}

	db.D.Storage.Delete(&p)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Delete successfully!"})
}

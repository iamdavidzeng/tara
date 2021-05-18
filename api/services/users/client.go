package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdavidzeng/tara/api/services"
	"github.com/iamdavidzeng/tara/internal/db"
)

type Users struct {
	services.GormBase
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
}

var UserOperator *Users = &Users{}

func (u Users) New(c *gin.Context) {
	c.BindJSON(&u)

	db.D.Storage.Create(&u)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Create successfully!"})
}

func (u Users) Get(c *gin.Context) {
	id := c.Param("id")

	db.D.Storage.First(&u, id)

	if u.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": u})
}

func (u Users) List(c *gin.Context) {
	users := []Users{}

	db.D.Storage.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func (u Users) Update(c *gin.Context) {
	id := c.Param("id")

	db.D.Storage.First(&u, id)
	if u.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	c.BindJSON(&u)

	db.D.Storage.Model(&u).Updates(u)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Update successfully!"})
}

func (u Users) Del(c *gin.Context) {
	id := c.Param("id")

	db.D.Storage.First(&u, id)
	if u.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	db.D.Storage.Delete(&u)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Delete successfully!"})
}

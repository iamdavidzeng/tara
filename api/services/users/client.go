package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdavidzeng/tara/api/models"
	"github.com/iamdavidzeng/tara/api/schemas"
	"github.com/iamdavidzeng/tara/internal/db"
)

func New(c *gin.Context) {
	var data schemas.UserSchema
	c.BindJSON(&data)

	_user := models.Users{Email: data.Email, Phone: data.Phone, Password: data.Password}
	db.D.Storage.Create(&_user)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Create successfully!"})
}

func Get(c *gin.Context) {
	var user models.Users
	userID := c.Param("id")

	db.D.Storage.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	_user := schemas.UserSchema{ID: user.ID, Email: user.Email, Phone: user.Phone, Password: user.Password}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _user})
}

func List(c *gin.Context) {
	var users []models.Users
	var data []schemas.UserSchema = []schemas.UserSchema{}

	db.D.Storage.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "data": data})
		return
	}

	for _, user := range users {
		data = append(data, schemas.UserSchema{ID: user.ID, Email: user.Email, Phone: user.Phone, Password: user.Password})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func Update(c *gin.Context) {
	var user models.Users
	userID := c.Param("id")

	db.D.Storage.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	var data schemas.UserSchema
	c.BindJSON(&data)

	db.D.Storage.Model(&user).Update("email", data.Email)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Update successfully!"})
}

func Delete(c *gin.Context) {
	var user models.Users
	userID := c.Param("id")

	db.D.Storage.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	db.D.Storage.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Delete successfully!"})
}

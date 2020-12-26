package services

import (
	"net/http"
	"tara/api/db"
	"tara/api/models"
	"tara/api/schemas"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user schemas.UserSchema
	c.BindJSON(&user)

	_user := models.Users{Email: user.Email, Phone: user.Phone, Password: user.Password}
	db.DB.Create(&_user)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Create successfully!"})
}

func GetUser(c *gin.Context) {
	var user models.Users
	userID := c.Param("id")

	db.DB.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	_user := schemas.UserSchema{ID: user.ID, Email: user.Email, Phone: user.Phone, Password: user.Password}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _user})
}

func GetUsers(c *gin.Context) {
	var users []models.Users
	var _users []schemas.UserSchema

	db.DB.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusUnauthorized, "data": "User Not Found!"})
		return
	}

	for _, user := range users {
		_users = append(_users, schemas.UserSchema{ID: user.ID, Email: user.Email, Phone: user.Phone, Password: user.Password})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _users})
}

func UpdateUser(c *gin.Context) {
	var user models.Users
	userID := c.Param("id")

	db.DB.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	var userData schemas.UserSchema
	c.BindJSON(&userData)

	db.DB.Model(&user).Update("email", userData.Email)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Update successfully!"})
}

func DeleteUser(c *gin.Context) {
	var user models.Users
	userID := c.Param("id")

	db.DB.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	db.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Delete successfully!"})
}

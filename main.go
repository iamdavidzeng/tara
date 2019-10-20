package main

import (
	"net/http"
	"os/user"
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// models: Users
type (
	Users struct {
		gorm.Model
		Email    string
		Phone    string
		Password string
	}
)

// Schemas: UserSchema
type (
	UserSchema struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
)

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@/demo?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Users{})
}

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "hello, world!"})
	})

	userAPI := router.Group("/api/v1/users")

	{
		userAPI.GET("/", getUsers)
		userAPI.GET("/:id", getUser)
		userAPI.POST("/", createUser)
		userAPI.POST("/:id", updateUser)
		userAPI.DELETE("/:id", deleteUser)
	}

	router.Run()
}

func createUser(c *gin.Context) {
	return
}

func getUser(c *gin.Context) {
	var user Users
	userID := c.Param("id")

	db.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	_user := UserSchema{ID: user.ID, Email: user.Email, Phone: user.Phone, Password: user.Password}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _user})
}

func getUsers(c *gin.Context) {
	return
}

func updateUser(c *gin.Context) {
	return
}

func deleteUser(c *gin.Context) {
	return
}

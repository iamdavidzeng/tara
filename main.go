package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
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
	db, err = gorm.Open("mysql", "root:@/demo?charset=utf8&parseTime=true")
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
	user := Users{Email: c.PostForm("email"), Phone: c.PostForm("phone"), Password: c.PostForm("password")}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Create successfully!"})
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
	var users []Users
	var _users []UserSchema

	db.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusUnauthorized, "data": "User Not Found!"})
		return
	}

	for _, user := range users {
		_users = append(_users, UserSchema{ID: user.ID, Email: user.Email, Phone: user.Phone, Password: user.Password})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _users})
}

func updateUser(c *gin.Context) {
	var user Users
	userID := c.Param("id")

	db.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	db.Model(&user).Update("email", c.PostForm("email"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Update successfully!"})
}

func deleteUser(c *gin.Context) {
	var user Users
	userID := c.Param("id")

	db.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "User Not Found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Delete successfully!"})
}

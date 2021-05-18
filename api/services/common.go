package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ModelOperator interface {
	New(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Del(c *gin.Context)
}

type GormBase struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

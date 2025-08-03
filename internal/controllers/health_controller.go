package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthController struct {
	db *gorm.DB
}

func NewHealthController(db *gorm.DB) *HealthController {
	return &HealthController{db: db}
}

func (c *HealthController) HealthCheck(ctx *gin.Context) {
	dbStatus := "connected"

	// Test database connection
	sqlDB, err := c.db.DB()
	if err != nil {
		dbStatus = "disconnected"
	} else {
		err = sqlDB.Ping()
		if err != nil {
			dbStatus = "disconnected"
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "available",
		"message": "API is running smoothly",
		"data": gin.H{
			"version":     "1.0.0",
			"environment": gin.Mode(),
			"database":    dbStatus,
		},
	})
}

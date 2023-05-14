package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Set up user routes
	userGroup := r.Group("/user")
	UserRoutes(userGroup)

	// TODO: Set up other routes

	return r
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nicenonecb/GoVersatile/internal/app/user/models"
	"net/http"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// TODO: Save the user to the database

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	})

	router.POST("/login", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// TODO: Verify the user's credentials

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})
}

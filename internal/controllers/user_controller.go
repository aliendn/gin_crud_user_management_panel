package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}

func Logout(c *gin.Context) {
	// Here you can handle token invalidation if needed
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

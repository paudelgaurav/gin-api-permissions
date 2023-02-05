package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-api-permissions/database"
	"github.com/paudelgaurav/gin-api-permissions/models"
	"github.com/paudelgaurav/gin-api-permissions/utils"
)

func BasicAuthPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the Basic Authentication credentials
		username, _, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		var user models.User
		if err := database.DB.First(&user, "username = ?", username).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		if !utils.Exists(permission, user.Permissions) {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

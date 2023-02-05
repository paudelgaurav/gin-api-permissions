package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-api-permissions/database"
	"github.com/paudelgaurav/gin-api-permissions/models"
)

func Ping(c *gin.Context) {
	var user models.User
	database.DB.First(&user)

	c.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}

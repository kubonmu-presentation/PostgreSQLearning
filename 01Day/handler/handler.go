package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health_handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"database": "connected",
	})
}

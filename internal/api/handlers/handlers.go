package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

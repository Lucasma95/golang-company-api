package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary This endpoint is used to verify if the api is running
// @ID health
// @Produce json
// @Success 200
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "running",
	})
}
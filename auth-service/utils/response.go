package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string      `json:"message" example:"Error"`
	Details interface{} `json:"details,omitempty"`
}

func ResponseOK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

func ResponseCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

func ResponseError(c *gin.Context, code int, message string, errors interface{}) {
	c.JSON(code, gin.H{
		"status":  false,
		"message": message,
		"errors":  errors,
	})
}

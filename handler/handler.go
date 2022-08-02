package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler - stores pointer to our comment service
type Config struct {
	R *gin.Engine
}

//create new handler which return a pointer to handler
type Handler struct {
}

//setup routes for all routes in application
func SetupRoutes(c *Config) {
	fmt.Println("Setting up routes")
	handler := &Handler{}
	router := c.R

	router.GET("/api/health", handler.ApiHealth)
}

func (h *Handler) ApiHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "i'm alive",
	})
}

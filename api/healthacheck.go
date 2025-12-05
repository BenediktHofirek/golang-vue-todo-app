package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) setupHealthcheck() {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})
}

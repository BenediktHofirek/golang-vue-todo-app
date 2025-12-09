package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code int
	Data any
	Error string
}

func ResponseOk(c *gin.Context, data any) {
  c.JSON(http.StatusOK, data)
}

func ResponseNoContent(c *gin.Context) {
  c.JSON(http.StatusNoContent, gin.H{})
}

func ResponseCreated(c *gin.Context, data any) {
  c.JSON(http.StatusCreated, data)
}

func ResponseBadRequest(c *gin.Context, err error) {
  c.JSON(http.StatusBadRequest, APIResponse{
    Code:    http.StatusBadRequest,
    Error:   err.Error(),
  })
}

func ResponseNotFound(c *gin.Context, err error) {
  c.JSON(http.StatusNotFound, APIResponse{
    Code:    http.StatusNotFound,
    Error:   err.Error(),
  })
}

func ResponseInternalServerError(c *gin.Context, err error) {
  c.JSON(http.StatusInternalServerError, APIResponse{
    Code:    http.StatusInternalServerError,
    Error:   err.Error(),
  })
}

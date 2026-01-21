package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/BenediktHofirek/golang-vue-todo-app/db"
	"github.com/gin-gonic/gin"
)

func (s *Server) getTodos(c *gin.Context) {
	token := GetUserAuthToken(c)

	todos, err := s.queries.Todo_GetMany(c.Request.Context(), token.UID)
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseOk(c, todos)
}

func (s *Server) getTodoById(c *gin.Context) {
	token := GetUserAuthToken(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo id: " + c.Param("id")})
		return
	}

	todo, err := s.queries.Todo_GetOneById(
		c.Request.Context(),
		db.Todo_GetOneByIdParams{
			ID:     id,
			UserID: token.UID,
		},
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ResponseNotFound(c, err)
		} else {
			ResponseInternalServerError(c, err)
		}

		return
	}

	ResponseOk(c, todo)
}

func (s *Server) deleteTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo id: " + c.Param("id")})
		return
	}

	token := GetUserAuthToken(c)

	if err := s.queries.Todo_DeleteOne(
		c.Request.Context(),
		db.Todo_DeleteOneParams{
			ID:     id,
			UserID: token.UID,
		},
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ResponseNotFound(c, err)
		} else {
			ResponseInternalServerError(c, err)
		}

		return
	}

	ResponseNoContent(c)
}

func (s *Server) createBlankTodo(c *gin.Context) {
	token := GetUserAuthToken(c)

	todoId, err := s.queries.Todo_CreateBlank(c.Request.Context(), token.UID)
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseCreated(c, todoId)
}

type UpdateTodoDto struct {
	Title         *string    `json:"title" binding:"omitempty"`
	Description   *string    `json:"description" binding:"omitempty"`
	Completed     *bool      `json:"completed" binding:"omitempty,boolean"`
	DueDate       *time.Time `json:"due_date" binding:"omitempty,gt=now"`
	Starred       *bool      `json:"starred" binding:"omitempty,boolean"`
	ScheduledDate *time.Time `json:"scheduled_date" binding:"omitempty"`
}

func (s *Server) updateTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo id: " + c.Param("id")})
		return
	}

	var dto UpdateTodoDto

	if err := c.ShouldBind(&dto); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	if dto.Title != nil {
		trimmedTitle := strings.TrimSpace(*dto.Title)
		dto.Title = &trimmedTitle
	}

	if dto.Description != nil {
		trimmedDescription := strings.TrimSpace(*dto.Description)
		dto.Description = &trimmedDescription
	}

	token := GetUserAuthToken(c)

	todo, err := s.queries.Todo_UpdateOne(c.Request.Context(), db.Todo_UpdateOneParams{
		ID:            id,
		UserID:        token.UID,
		Title:         dto.Title,
		Description:   dto.Description,
		Completed:     dto.Completed,
		DueDate:       toTimestamptz(dto.DueDate),
		Starred:       dto.Starred,
		ScheduledDate: toTimestamptz(dto.ScheduledDate),
	})

	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseOk(c, todo)
}

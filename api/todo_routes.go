package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/BenediktHofirek/golang-vue-todo-app/db"
	"github.com/gin-gonic/gin"
)

func (s *Server) getTodos(c *gin.Context) {
	token := GetUserAuthToken(c)

	todos, err := s.queries.Todo_GetManyByUserId(c.Request.Context(), token.UID)
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

type CreateTodoDto struct {
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed" binding:"boolean"`
	DueDate     *time.Time `json:"due_date" binding:"gt=now"`
}

func (s *Server) createTodo(c *gin.Context) {
	var dto CreateTodoDto

	if err := c.ShouldBind(&dto); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	token := GetUserAuthToken(c)

	todo, err := s.queries.Todo_CreateOne(c.Request.Context(), db.Todo_CreateOneParams{
		UserID:      token.UID,
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   dto.Completed,
		DueDate:     toTimestamptz(dto.DueDate),
	})
	if err != nil {
		ResponseInternalServerError(c, err)
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

type UpdateTodoDto struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Completed   *bool      `json:"completed" binding:"omitempty,boolean"`
	DueDate     *time.Time `json:"due_date" binding:"omitempty,gt=now"`
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

	token := GetUserAuthToken(c)

	todo, err := s.queries.Todo_UpdateOne(c.Request.Context(), db.Todo_UpdateOneParams{
		ID:          id,
		UserID:      token.UID,
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   dto.Completed,
		DueDate:     toTimestamptz(dto.DueDate),
	})

	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseOk(c, todo)
}

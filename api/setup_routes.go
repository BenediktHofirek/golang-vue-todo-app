package main

func (s *Server) setupRoutes() {
	v1 := s.router.Group("/api/v1")
	v1.Use(AuthMiddleware())

	{
		v1.GET("/todos", s.getTodos)
		v1.GET("/todos/:id", s.getTodoById)
		v1.POST("/todos", s.createTodo)
		v1.PATCH("/todos/:id", s.updateTodo)
		v1.DELETE("/todos/:id", s.deleteTodo)
	}
}

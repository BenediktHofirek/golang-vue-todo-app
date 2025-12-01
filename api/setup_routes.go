package main

func (s *Server) setupRoutes() {
	v1 := s.router.Group("/api/v1")
	{
		v1.GET("/todos/:userId", s.getTodos)
		v1.GET("/todos/:userId/:id", s.getTodoById)
		v1.POST("/todos", s.createTodo)
		v1.PATCH("/todos/:userId/:id", s.updateTodo)
		v1.DELETE("/todos/:userId/:id", s.deleteTodo)
	}
}

package main

import (
	"log"
	"os"

	"github.com/BenediktHofirek/golang-vue-todo-app/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func main() {
	database, err := db.CreateDbPool()
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	defer database.Close()

	queries := db.New(database)

	router := gin.Default()

	server := &Server{
		queries: queries,
		router:  router,
	}

	server.setupHealthcheck()
	server.setupRoutes()

	port := os.Getenv("API_PORT")
	server.router.Run("0.0.0.0:" + port)
}

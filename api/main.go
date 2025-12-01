package main

import (
	"log"
	"os"

	"github.com/BenediktHofirek/golang-vue-todo-app/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	queries *db.Queries
	router *gin.Engine
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.CreateDbPool()
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	defer database.Close()

	queries := db.New(database);

	router := gin.Default()

	server := &Server{ 
		queries: queries,
		router: router,
	}

	server.setupRoutes()

	port := os.Getenv("PORT")
	server.router.Run(":" + port)
}

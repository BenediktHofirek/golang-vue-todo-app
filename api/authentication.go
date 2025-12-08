package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

const UserAuthTokenKey = "FIREBASE_USER_AUTH_TOKEN"

func AuthMiddleware() gin.HandlerFunc {
	projectId := os.Getenv("PROJECT_ID")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: projectId,
	})

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		if idToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing or malformed"})
			c.Abort()
			return
		}

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token", "details": err.Error()})
			c.Abort()
			return
		}

		c.Set(UserAuthTokenKey, token)

		c.Next()
	}
}

func GetUserAuthToken(c *gin.Context) *auth.Token {
	claims := c.MustGet(UserAuthTokenKey).(*auth.Token)
	return claims
}

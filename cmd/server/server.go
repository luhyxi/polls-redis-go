package server

import (
	"log"
	"net/http"

	"example.com/go-polls/internal"
	"example.com/go-polls/pkg/models"
	"example.com/go-polls/pkg/services/user"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "go-polls",
		})
	})

	// User endpoints
	router.POST("/users", func(ctx *gin.Context) {
		var req models.CreateUserRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		user, err := user.CreateUser(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "user created successfully",
			"user":    user,
		})
	})

	// Get user by ID
	router.GET("/users/:id", func(ctx *gin.Context) {
		userID := ctx.Param("id")
		
		user, err := user.GetUser("user:" + userID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	// Login endpoint
	router.POST("/login", func(ctx *gin.Context) {
		var loginReq struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&loginReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		// Simple validation - It accepts any password
		_, err := user.GetUser("user:" + loginReq.Username)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// Generate JWT token
		token, err := internal.GenerateJwtToken(loginReq.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"token":   token,
		})
	})

	// Get all users
	router.GET("/users/", func(ctx *gin.Context) {
		user, err := user.GetAllUsers()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "users not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	// Start server on port 8080
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

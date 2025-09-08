package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	poll "example.com/go-polls/pkg/services/poll"
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
		var req struct {
			Name string `json:"name" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		user, err := poll.CreateUser(req.Name)
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
		
		user, err := poll.GetUser("user:" + userID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
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

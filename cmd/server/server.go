package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	poll "example.com/go-polls/pkg/services/poll"
)

func main() {
	router := gin.Default()

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

	router.Run()
}

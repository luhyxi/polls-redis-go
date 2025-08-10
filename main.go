package main

import (
	"net/http"

	"example.com/go-polls/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		redisURL, err := internal.GetRedisURL()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error: %s", err.Error())
			return
		}
		ctx.String(http.StatusOK, redisURL)
	})

	router.Run()
}

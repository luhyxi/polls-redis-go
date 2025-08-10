package main

import (
    "net/http"
	"example.com/umamusume/pkg/services"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/albums", services.getAlbums)

    router.Run("localhost:8080")
}

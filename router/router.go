package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router
	router := gin.Default()

	// Initialize routes
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Run the server
	router.Run(":8080")
}

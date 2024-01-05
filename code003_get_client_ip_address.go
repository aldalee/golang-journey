package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		// obtain the IP address of the client
		clientIP := ctx.ClientIP()
		ctx.String(http.StatusOK, "Client IP: %s", clientIP)
	})
	// start the Gin server and listen on port 8080
	r.Run(":8080")
}

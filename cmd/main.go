package main

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/package/http"
)

func main() {
	router := gin.New()

	v1 := router.Group("/v1")
	{
		v1.POST("/create", http.Create )
		//v1.POST("/read", submitEndpoint)
		//v1.POST("/update", readEndpoint)
	}
	router.Run(":8080")
}

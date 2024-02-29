package main

import (
	deploy "github.com/NicolasLopes7/shipthing/lib/deploy"
	"github.com/gin-gonic/gin"
)

func InitHTTPServer() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	r.POST("/deploy", deploy.Handler)

	r.Run(":3000")
}

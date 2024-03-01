package main

import (
	config "github.com/NicolasLopes7/shipthing/config"
	"github.com/NicolasLopes7/shipthing/uploader-service/deploy"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitConfig()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	r.POST("/deploy", deploy.Handler)

	r.Run(":3000")
}

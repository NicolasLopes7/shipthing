package deploy

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	fs "github.com/NicolasLopes7/shipthing/lib/fs"
	github "github.com/NicolasLopes7/shipthing/lib/github"
)

func Handler(ctx *gin.Context) {
	var req *DeployRequest
	err := ctx.BindJSON(&req)

	if err != nil || req.URL == "" {
		ctx.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	deployId := uuid.New()
	path := fmt.Sprintf("/tmp/%s", deployId)
	_, err = github.Clone(req.URL, path)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = fs.WalkDir(path, func(path string) error {
		time.Sleep(5 * time.Second)
		return nil
	})

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"id": deployId,
	})
}

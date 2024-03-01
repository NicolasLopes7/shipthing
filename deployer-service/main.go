package main

import (
	"context"
	"fmt"

	"github.com/NicolasLopes7/shipthing/config"
	aws "github.com/NicolasLopes7/shipthing/lib/aws"
)

func main() {
	err := config.InitConfig()

	redisCtx := context.Background()

	if err != nil {
		panic(err)
	}

	for true {
		res := config.RedisClient.BRPop(redisCtx, 0, "builds")

		if res.Err() != nil {
			fmt.Println(res.Err())
			continue
		}

		deployId := res.Val()[1]
		fmt.Println("deploying: ", deployId)
		err := aws.DownloadS3Folder(deployId)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("downloaded artifacts from s3: ", deployId)

	}
}

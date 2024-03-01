package main

import (
	"context"
	"fmt"

	"github.com/NicolasLopes7/shipthing/config"
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

		fmt.Println("Deploying", deployId)
	}
}

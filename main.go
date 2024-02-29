package main

import (
	config "github.com/NicolasLopes7/shipthing/config"
)

func main() {
	err := config.InitConfig()

	if err != nil {
		panic(err)
	}

	InitHTTPServer()
}

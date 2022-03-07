package main

import (
	"log"
	"simple-api-gateway/config"
	"simple-api-gateway/pkg/api"
)

func main() {
	api.Start()
}

func init() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("error load env")
	}
}

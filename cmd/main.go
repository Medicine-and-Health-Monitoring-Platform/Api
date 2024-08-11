package main

import (
	"Api_Gateway/api"
	"Api_Gateway/config"
)

func main() {
	cfg := config.Load()

	router := api.NewRouter(cfg)

	router.Run(cfg.HTTP_PORT)
}

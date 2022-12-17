package main

import (
	"log"
	"yarus-tz/config"
	"yarus-tz/internal/app"
)

func main() {

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}

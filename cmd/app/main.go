package main

import (
	"log"

	"github.com/vasolovev/secret_santa/config"
	"github.com/vasolovev/secret_santa/internal/app"
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

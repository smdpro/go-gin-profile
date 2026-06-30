package main

import (
	"go-gin-profile/internal/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config:", err)
		return
	}
	log.Println("Config loaded successfully:")
	log.Println("MongoURI:", cfg.MongoURI)
	log.Println("Port:", cfg.Port)
}

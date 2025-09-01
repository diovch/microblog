package main

import (
	"log"

	"github.com/diovch/microblog/config"
	"github.com/diovch/microblog/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
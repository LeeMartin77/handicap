package main

import (
	"context"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/leemartin77/handicap/internal/config"
	"github.com/leemartin77/handicap/internal/server"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	ctx := context.Background()

	cfg := config.Config{}

	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal(err)
	}
	srv, err := server.NewServer(ctx, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.RunServer(); err != nil {
		log.Fatal(err)
	}
	log.Println("server shut down")
}

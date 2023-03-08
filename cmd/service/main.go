package main

import (
	"context"
	"log"

	"github.com/IgorAleksandroff/agent-status/internal/config"
	"github.com/IgorAleksandroff/agent-status/internal/repository"
	"github.com/IgorAleksandroff/agent-status/internal/service"
	"github.com/IgorAleksandroff/agent-status/internal/usecase"
)

func main() {
	log.Println("debug: start main")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.GetConfig()

	repo, err := repository.NewPostgres(ctx, cfg.DataBaseURI)
	if err != nil {
		log.Fatalf("failed to create repository: %s", err)
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Printf("failed to close repository: %s", err)
		}
	}()

	auth := usecase.NewAuthorization(repo)
	statusUC := usecase.NewStatus(repo)

	app, err := service.New(cfg.ServerConfig, auth, statusUC)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}

	app.Run(ctx)
}

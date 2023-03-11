package main

import (
	"context"
	"log"

	"github.com/IgorAleksandroff/agent-status/internal/config"
	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/repository"
	"github.com/IgorAleksandroff/agent-status/internal/service"
	"github.com/IgorAleksandroff/agent-status/internal/usecase"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command"
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
	// для прода надо использовать внешний брокер, который поддерживает несколько инстансов
	commandQueue := make(chan entity.Command, cfg.QueueSize)
	commandSender := external_command.NewSender(commandQueue)
	// todo executor
	statusUC := usecase.NewStatus(repo, commandSender)

	app, err := service.New(cfg.ServerConfig, auth, statusUC)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}

	app.Run(ctx)
}

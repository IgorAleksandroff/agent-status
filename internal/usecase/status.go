package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command/commands"
	"github.com/pkg/errors"
)

const maxRetries = 2

//go:generate mockery --name Status --with-expecter
//go:generate mockery --name statusRepository --with-expecter
//go:generate mockery --name statusSender --with-expecter

type statusUsecase struct {
	repo            statusRepository
	externalCommand statusSender
}

type statusSender interface {
	Send(cmd external_command.Base) error
}

type Status interface {
	AgentSetStatus(ctx context.Context, agent entity.Agent, mode entity.Mode) error
}

type statusRepository interface {
	AgentSetStatusTx(ctx context.Context, agent entity.Agent, mode entity.Mode) (*int64, error)
}

func NewStatus(r statusRepository, s statusSender) *statusUsecase {
	return &statusUsecase{
		repo:            r,
		externalCommand: s,
	}
}

func (s statusUsecase) AgentSetStatus(ctx context.Context, agent entity.Agent, mode entity.Mode) error {
	_, err := s.repo.AgentSetStatusTx(ctx, agent, mode)
	if err != nil {
		return errors.WithStack(err)
	}

	// отправка сообщений о начале и завершении смены, порядок не важен
	cmdMsg, err := commands.NewSendMessage(map[string]string{
		"login":     agent.Login,
		"status":    string(*agent.Status),
		"changedAt": time.Now().Format(time.RFC3339),
		"counter":   strconv.FormatInt(maxRetries, 10),
	})
	if err != nil {
		log.Println("failed to create command: send message")
	}

	err = s.externalCommand.Send(cmdMsg)
	if err != nil {
		log.Printf("failed to send command: %+v", cmdMsg)
	}

	// todo отправка статусов в сервис автоназначения писем и диалогов
	// порядок важен, либо синхронно в транзакции,
	// либо асинхронно со строгим порядком или с проверкой логов, не отправлять если состояние изменилось

	return nil
}

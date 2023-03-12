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
	logID, err := s.repo.AgentSetStatusTx(ctx, agent, mode)
	if err != nil {
		return errors.WithStack(err)
	}

	// отправка сообщений о начале и завершении смены, порядок не важен
	if *agent.Status == entity.StatusActive || *agent.Status == entity.StatusInactive {
		s.sendMessage(agent)
	}

	// отправка статусов в сервис автоназначения писем и диалогов
	// с проверкой логов, не отправляется, если статус изменился
	if logID != nil {
		s.sendToAutoAssignment(agent, *logID)
	}

	return nil
}

func (s statusUsecase) sendMessage(agent entity.Agent) {
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
}

func (s statusUsecase) sendToAutoAssignment(agent entity.Agent, logID int64) {
	cmdMsg, err := commands.NewSendMessage(map[string]string{
		"login":   agent.Login,
		"status":  string(*agent.Status),
		"logID":   strconv.FormatInt(logID, 10),
		"counter": strconv.FormatInt(maxRetries, 10),
	})
	if err != nil {
		log.Println("failed to create command: update auto-assignment")
	}

	err = s.externalCommand.Send(cmdMsg)
	if err != nil {
		log.Printf("failed to send command: %+v", cmdMsg)
	}
}

package usecase

import (
	"context"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command"
)

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
	// todo отправка сообщений о начале и завершении смены
	// порядок не важен, можно отправить в очередь после успешной транзакции

	// todo отправка статусов в сервис автоназначения писем и диалогов
	// порядок важен, либо синхронно в транзакции,
	// либо асинхронно со строгим порядком или с проверкой логов, не отправлять если состояние изменилось

	_, err := s.repo.AgentSetStatusTx(ctx, agent, mode)

	return err
}

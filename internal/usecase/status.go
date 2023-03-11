package usecase

import (
	"context"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
)

//go:generate mockery --name Status --with-expecter
//go:generate mockery --name statusRepository --with-expecter

type statusUsecase struct {
	repo statusRepository
}

type Status interface {
	AgentSetStatus(ctx context.Context, agent entity.Agent, mode entity.Mode) error
}

type statusRepository interface {
	AgentSetStatusTx(ctx context.Context, agent entity.Agent, mode entity.Mode) (*int64, error)
}

func NewStatus(r statusRepository) *statusUsecase {
	return &statusUsecase{repo: r}
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

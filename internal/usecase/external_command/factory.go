package external_command

import (
	"fmt"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command/commands"
	"github.com/pkg/errors"
)

type factory struct {
	repo commands.AutoAssignmentRepository
}

func NewFactory(repo commands.AutoAssignmentRepository) *factory {
	return &factory{
		repo: repo,
	}
}

func (f factory) GetCommandFromType(commandType entity.CommandType, params map[string]string) (Base, Executor, error) {
	switch commandType {
	case entity.SendMsg:
		e := commands.NewSendMessageExecutor()
		c, err := commands.NewSendMessage(params)
		if err != nil {
			return nil, nil, errors.Wrap(err, fmt.Sprintf("failed to create command: %s", entity.SendMsg))
		}
		return c, e, nil

	case entity.AutoAssignment:
		e := commands.NewSendToAutoAssignmentExecutor(f.repo)
		c, err := commands.NewSendToAutoAssignment(params)
		if err != nil {
			return nil, nil, errors.Wrap(err, fmt.Sprintf("failed to create command: %s", entity.AutoAssignment))
		}
		return c, e, nil

	default:
		return nil, nil, errors.Errorf("not found suitable command: %s", commandType)
	}
}

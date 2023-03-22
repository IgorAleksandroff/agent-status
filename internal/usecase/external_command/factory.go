package external_command

import (
	"fmt"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command/commands"
)

type factory struct {
	repo         commands.AutoAssignmentRepository
	messenger    commands.MessageSender
	statusSender commands.StatusSender
}

func NewFactory(msg commands.MessageSender, repo commands.AutoAssignmentRepository, statusSender commands.StatusSender) *factory {
	return &factory{
		repo:         repo,
		messenger:    msg,
		statusSender: statusSender,
	}
}

func (f factory) GetCommandFromType(commandType entity.CommandType, params map[string]string) (commands.Base, commands.Executor, error) {
	switch commandType {
	case entity.SendMsg:
		e := commands.NewSendMessageExecutor(f.messenger)
		c, err := commands.NewSendMessage(params)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create command- %s: %w", entity.SendMsg, err)
		}
		return c, e, nil

	case entity.AutoAssignment:
		e := commands.NewSendToAutoAssignmentExecutor(f.repo, f.statusSender)
		c, err := commands.NewSendToAutoAssignment(params)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create command- %s: %w", entity.AutoAssignment, err)
		}
		return c, e, nil

	default:
		return nil, nil, fmt.Errorf("not found suitable command: %s", commandType)
	}
}

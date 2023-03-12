package external_command

import (
	"fmt"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command/commands"
	"github.com/pkg/errors"
)

type factory struct {
}

func NewFactory() *factory {
	return &factory{}
}

func (f factory) GetCommandFromType(commandType entity.CommandType, params map[string]string) (Base, Executor, error) {
	switch commandType {
	case entity.SendMsg:
		c, err := commands.NewSendMessage(params)
		if err != nil {
			return nil, nil, errors.Wrap(err, fmt.Sprintf("failed to create command: %s", entity.SendMsg))
		}

		e := commands.NewSendMessageExecutor()

		return c, e, nil
	default:
		return nil, nil, errors.Errorf("not found suitable command: %s", commandType)
	}
}

package external_command

import (
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command/commands"
	"github.com/pkg/errors"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
)

type sender struct {
	queue chan entity.Event
}

func NewSender(q chan entity.Event) *sender {
	return &sender{
		queue: q,
	}
}

func (s sender) Send(c commands.Base) error {
	if c == nil {
		return errors.New("command is nil")
	}

	s.queue <- entity.Event{
		Type:   c.Type(),
		Params: c.Params(),
	}

	return nil
}

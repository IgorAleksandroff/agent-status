package external_command

import (
	"github.com/pkg/errors"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
)

type sender struct {
	queue chan entity.Command
}

func NewSender(q chan entity.Command) sender {
	return sender{
		queue: q,
	}
}

func (s sender) Send(c Base) error {
	if c == nil {
		return errors.New("command is nil")
	}

	s.queue <- entity.Command{
		Type:   c.Type(),
		Params: c.Params(),
	}

	return nil
}

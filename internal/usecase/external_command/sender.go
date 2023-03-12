package external_command

import (
	"github.com/pkg/errors"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
)

type Base interface {
	Type() entity.CommandType
	Params() *map[string]string
}

type sender struct {
	queue chan entity.Event
}

func NewSender(q chan entity.Event) *sender {
	return &sender{
		queue: q,
	}
}

func (s sender) Send(c Base) error {
	if c == nil {
		return errors.New("command is nil")
	}

	s.queue <- entity.Event{
		Type:   c.Type(),
		Params: c.Params(),
	}

	return nil
}

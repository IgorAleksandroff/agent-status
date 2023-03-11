package commands

import (
	"context"
	"reflect"

	"github.com/pkg/errors"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
)

type Base interface {
	Type() entity.CommandType
	ExecutorType() reflect.Type
	Params() *map[string]string
}

type Executor interface {
	ValidityCheck(ctx context.Context, command Base) bool
	Execute(ctx context.Context, command Base) error
	Retry(ctx context.Context, command Base) bool
}

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

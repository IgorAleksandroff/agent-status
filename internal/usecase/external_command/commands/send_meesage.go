package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/pkg/errors"
)

var (
	_ Base     = (*sendMessage)(nil)
	_ Executor = (*sendMessageExecutor)(nil)
)

type MessageSender interface {
	Send(msg string) error
}

type sendMessage struct {
	Agent     string
	Status    string
	ChangedAt string
	counter   string
}

type sendMessageExecutor struct {
	client MessageSender
}

func NewSendMessage(params map[string]string) (*sendMessage, error) {
	agent, ok := params["login"]
	if !ok {
		return nil, errors.New("missing parameter: login")
	}

	status, ok := params["status"]
	if !ok {
		return nil, errors.New("missing parameter: status")
	}

	changedAt, ok := params["changedAt"]
	if !ok {
		return nil, errors.New("missing parameter: changedAt")
	}

	counter, ok := params["counter"]
	if !ok {
		return nil, errors.New("missing parameter: counter")
	}

	return &sendMessage{
		Agent:     agent,
		Status:    status,
		ChangedAt: changedAt,
		counter:   counter,
	}, nil
}

func (c sendMessage) Type() entity.CommandType {
	return entity.SendMsg
}

func (c sendMessage) Params() *map[string]string {
	return &map[string]string{
		"login":     c.Agent,
		"status":    c.Status,
		"changedAt": c.ChangedAt,
		"counter":   c.counter,
	}
}

func NewSendMessageExecutor(client MessageSender) *sendMessageExecutor {
	return &sendMessageExecutor{
		client: client,
	}
}

func (e sendMessageExecutor) ValidityCheck(ctx context.Context, command Base) bool {
	return true
}

func (e sendMessageExecutor) Execute(ctx context.Context, command Base) error {
	p := *command.Params()

	return e.client.Send(fmt.Sprintf(
		"Agent %s change status to %s at %s.",
		p["login"],
		p["status"],
		p["changedAt"],
	))
}

func (e sendMessageExecutor) Retry(ctx context.Context, command Base) bool {
	p := *command.Params()
	counter, ok := p["counter"]
	if !ok {
		return false
	}
	if c, err := strconv.Atoi(counter); err != nil || c > 0 {
		return true
	}

	return false
}

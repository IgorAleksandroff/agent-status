package commands

import (
	"context"
	"log"
	"strconv"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/pkg/errors"
)

var (
	_ Base     = (*sendMessage)(nil)
	_ Executor = (*sendMessageExecutor)(nil)
)

type AutoAssignmentRepository interface {
	GetLastLogIdForAgent(ctx context.Context, login string) (int64, error)
}

type StatusSender interface {
	Send(ctx context.Context, login, status string) error
}

type sendToAutoAssignment struct {
	Agent   string
	Status  string
	LogID   string
	counter string
}

type sendToAutoAssignmentExecutor struct {
	repo   AutoAssignmentRepository
	client StatusSender
}

func NewSendToAutoAssignment(params map[string]string) (*sendToAutoAssignment, error) {
	agent, ok := params["login"]
	if !ok {
		return nil, errors.New("missing parameter: login")
	}

	status, ok := params["status"]
	if !ok {
		return nil, errors.New("missing parameter: status")
	}

	logID, ok := params["logID"]
	if !ok {
		return nil, errors.New("missing parameter: changedAt")
	}

	counter, ok := params["counter"]
	if !ok {
		return nil, errors.New("missing parameter: counter")
	}

	return &sendToAutoAssignment{
		Agent:   agent,
		Status:  status,
		LogID:   logID,
		counter: counter,
	}, nil
}

func (c sendToAutoAssignment) Type() entity.CommandType {
	return entity.AutoAssignment
}

func (c sendToAutoAssignment) Params() *map[string]string {
	return &map[string]string{
		"login":   c.Agent,
		"status":  c.Status,
		"logID":   c.LogID,
		"counter": c.counter,
	}
}

func NewSendToAutoAssignmentExecutor(repo AutoAssignmentRepository, client StatusSender) *sendToAutoAssignmentExecutor {
	return &sendToAutoAssignmentExecutor{
		repo:   repo,
		client: client,
	}
}

func (e sendToAutoAssignmentExecutor) ValidityCheck(ctx context.Context, command Base) bool {
	actualLogID, err := e.repo.GetLastLogIdForAgent(ctx, (*command.Params())["login"])
	if err != nil {
		log.Printf("failed to GetLastLogIdForAgent: %s", (*command.Params())["login"])
		return false
	}

	if strconv.FormatInt(actualLogID, 10) == (*command.Params())["logID"] {
		return true
	}

	return false
}

func (e sendToAutoAssignmentExecutor) Execute(ctx context.Context, command Base) error {
	p := *command.Params()

	return e.client.Send(ctx, p["login"], p["status"])
}

func (e sendToAutoAssignmentExecutor) Retry(ctx context.Context, command Base) bool {
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

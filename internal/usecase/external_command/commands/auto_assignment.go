package commands

import (
	"context"
	"log"
	"strconv"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command"
	"github.com/pkg/errors"
)

var (
	_ external_command.Base     = (*sendMessage)(nil)
	_ external_command.Executor = (*sendMessageExecutor)(nil)
)

type AutoAssignmentRepository interface {
	GetLastLogIdForAgent(ctx context.Context, login string) (*int64, error)
}

type sendToAutoAssignment struct {
	Agent   string
	Status  string
	LogID   string
	counter string
}

type sendToAutoAssignmentExecutor struct {
	repo   AutoAssignmentRepository
	client string
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

func NewSendToAutoAssignmentExecutor(repo AutoAssignmentRepository) *sendToAutoAssignmentExecutor {
	return &sendToAutoAssignmentExecutor{
		repo:   repo,
		client: "to be done",
	}
}

func (e sendToAutoAssignmentExecutor) ValidityCheck(ctx context.Context, command external_command.Base) bool {
	actualLogID, err := e.repo.GetLastLogIdForAgent(ctx, (*command.Params())["login"])
	if err != nil {
		log.Printf("failed to GetLastLogIdForAgent: %s", (*command.Params())["login"])
		return false
	}

	if actualLogID != nil && strconv.FormatInt(*actualLogID, 10) == (*command.Params())["logID"] {
		return true
	}

	return false
}

func (e sendToAutoAssignmentExecutor) Execute(ctx context.Context, command external_command.Base) error {
	// todo:

	return nil
}

func (e sendToAutoAssignmentExecutor) Retry(ctx context.Context, command external_command.Base) bool {
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

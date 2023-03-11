package external_command

import (
	"context"
	"reflect"

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

package grpc

import "github.com/IgorAleksandroff/agent-status/internal/entity"

var grpcToStatus = []entity.Status{
	entity.StatusActive,
	entity.StatusReqInactive,
	entity.StatusInactive,
	entity.StatusReqBreak,
	entity.StatusBreak,
	entity.StatusForceMajeure,
	entity.StatusChat,
	entity.StatusLetter,
}

package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/generated/rpc"
	"github.com/IgorAleksandroff/agent-status/internal/usecase"
)

type handler struct {
	rpc.UnimplementedAgentStatusServer
	statusUC usecase.Status
}

func New(statusUC usecase.Status) *handler {
	return &handler{
		statusUC: statusUC,
	}
}

func (h handler) UserSetStatus(ctx context.Context, r *rpc.UserRequest) (*rpc.UserSetStatusResponse, error) {
	if r == nil || r.User == nil {
		return nil, status.Error(codes.Aborted, "empty user")
	}

	statusID := r.User.Status
	if int(statusID) >= len(grpcToStatus) {
		return nil, status.Errorf(codes.Aborted, "invalid status: %v", statusID)
	}

	if err := h.statusUC.AgentSetStatus(ctx, entity.Agent{
		Login:  r.User.Login,
		Status: &grpcToStatus[statusID],
	}, entity.Grpc); err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Internal, "failed to set status for %s", r.User.Login)
	}

	return &rpc.UserSetStatusResponse{}, nil
}

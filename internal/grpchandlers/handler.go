package grpchandlers

import (
	"context"
	"log"
	"time"

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

func (h handler) UserGetLog(ctx context.Context, r *rpc.UserRequest) (*rpc.UserGetLogResponse, error) {
	if r == nil || r.User == nil {
		return nil, status.Error(codes.Aborted, "empty user")
	}

	logs, err := h.statusUC.GetLogsForAgent(ctx, r.User.Login);
	if err != nil {
		log.Println(err.Error())

		return nil, status.Errorf(codes.Internal, "failed to set status for %s", r.User.Login)
	}

	result := make([]*rpc.Log, len(logs))
	for _, l := range logs {
		srcID, err := GetStatusIDbyName(string(l.Src))
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		dstID, err := GetStatusIDbyName(string(l.Dst))
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		result = append(result, &rpc.Log{
			Login:       l.Agent,
			Src:         rpc.Log_Statuses(srcID),
			Dst:         rpc.Log_Statuses(dstID),
			Mode:        rpc.Log_Modes(rpc.Log_Modes_value[string(l.Mode)]),
			ProcessedAt: l.ProcessedAt.Format(time.RFC3339),
		})
	}

	return &rpc.UserGetLogResponse{Logs: result}, nil
}

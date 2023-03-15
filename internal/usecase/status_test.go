package usecase_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command/commands"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/mocks"
)

func Test_statusUsecase_AgentSetStatus(t *testing.T) {
	type fields struct {
		repo            usecase.StatusRepository
		externalCommand usecase.StatusSender
	}
	type args struct {
		ctx   context.Context
		agent entity.Agent
		mode  entity.Mode
	}

	entityStatusInactive := entity.StatusInactive

	ctx := context.Background()
	agent := entity.Agent{
		Login:  "test",
		Status: &entityStatusInactive,
	}
	mode := entity.Rest
	msg, err := commands.NewSendMessage(map[string]string{
		"login":     agent.Login,
		"status":    string(*agent.Status),
		"changedAt": time.Now().Format(time.RFC3339),
		"counter":   strconv.FormatInt(2, 10),
	})
	if err != nil {
		t.Errorf("failed to create command: send message")
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success case",
			fields: fields{
				repo: func() usecase.StatusRepository {
					mockRepo := &mocks.StatusRepository{}
					mockRepo.EXPECT().AgentSetStatusTx(ctx, agent, mode).Return(nil, nil)

					return mockRepo
				}(),
				externalCommand: func() usecase.StatusSender {
					mockSender := &mocks.StatusSender{}
					mockSender.EXPECT().Send(msg).Return(nil)

					return mockSender
				}(),
			},
			args: args{
				ctx: ctx, agent: agent, mode: mode,
			},
			wantErr: false,
		},
		{
			name: "error case: repo returns error",
			fields: fields{
				repo: func() usecase.StatusRepository {
					mockRepo := &mocks.StatusRepository{}
					mockRepo.EXPECT().AgentSetStatusTx(ctx, agent, mode).Return(nil, errors.New("repo error"))

					return mockRepo
				}(),
			},
			args: args{
				ctx: ctx, agent: agent, mode: mode,
			},
			wantErr: true,
		},
		{
			name: "error case: sender returns error",
			fields: fields{
				repo: func() usecase.StatusRepository {
					mockRepo := &mocks.StatusRepository{}
					mockRepo.EXPECT().AgentSetStatusTx(ctx, agent, mode).Return(nil, nil)

					return mockRepo
				}(),
				externalCommand: func() usecase.StatusSender {
					mockSender := &mocks.StatusSender{}
					mockSender.EXPECT().Send(msg).Return(errors.New("sender error"))

					return mockSender
				}(),
			},
			args: args{
				ctx: ctx, agent: agent, mode: mode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := usecase.NewStatus(tt.fields.repo, tt.fields.externalCommand)

			if err := s.AgentSetStatus(tt.args.ctx, tt.args.agent, tt.args.mode); (err != nil) != tt.wantErr {
				t.Errorf("AgentSetStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

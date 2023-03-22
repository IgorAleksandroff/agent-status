package clients

import (
	"context"
	"log"

	"github.com/IgorAleksandroff/agent-status/internal/generated/rpc/clients"
	"github.com/IgorAleksandroff/agent-status/internal/grpchandlers"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clientGRPC struct {
	client clients.AutoAssigmentClient
	conn   *grpc.ClientConn
}

func NewGRPCAutoAssigment(socket string) (*clientGRPC, error) {
	conn, err := grpc.Dial(socket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	return &clientGRPC{
		client: clients.NewAutoAssigmentClient(conn),
		conn:   conn,
	}, nil
}

func (c clientGRPC) Send(ctx context.Context, login, status string) error {
	id, err := grpchandlers.GetStatusIDbyName(login)
	if err != nil {
		return err
	}

	_, err = c.client.UserChangeStatus(ctx, &clients.ChangeRequest{
		User: &clients.Agent{
			Login:  login,
			Status: clients.Agent_Statuses(id),
		},
	})

	return err
}

func (c clientGRPC) Close() {
	if err := c.conn.Close(); err != nil {
		log.Println(err)
	}
}

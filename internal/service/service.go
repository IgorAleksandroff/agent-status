package service

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 3 * time.Second
)

type Server interface {
	Run(ctx context.Context)
}

type server struct {
	serverHTTP   *http.Server
	serverGRPC   *grpc.Server
	gRPCListener net.Listener
}

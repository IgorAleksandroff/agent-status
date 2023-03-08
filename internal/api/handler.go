package api

import (
	"net/http"

	"github.com/IgorAleksandroff/agent-status/internal/usecase"
)

type handler struct {
	auth     usecase.Authorization
	statusUC usecase.Status
}

type handlerFunc interface {
	MethodFunc(method, path string, handler http.HandlerFunc)
}

func New(
	auth usecase.Authorization,
	status usecase.Status,
) *handler {
	return &handler{
		auth:     auth,
		statusUC: status,
	}
}

func (h *handler) Register(router handlerFunc, method, path string, handler http.HandlerFunc) {
	router.MethodFunc(method, path, handler)
}

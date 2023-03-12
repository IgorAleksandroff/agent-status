package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
)

func (h *handler) UserSetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contentTypeHeaderValue := r.Header.Get("Content-Type")
	if !strings.Contains(contentTypeHeaderValue, "application/json") {
		log.Println("unknown content-type")
		http.Error(w, "unknown content-type", http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		log.Println("empty body")
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	agent := entity.Agent{}
	reader := json.NewDecoder(r.Body)
	if err := reader.Decode(&agent); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.statusUC.AgentSetStatus(ctx, agent, entity.Rest); err != nil {
		log.Println(err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

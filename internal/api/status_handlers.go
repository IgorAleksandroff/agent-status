package api

import (
	"bytes"
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

func (h *handler) UserGetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	agent := entity.Agent{}
	if r.Body == nil {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	contentTypeHeaderValue := r.Header.Get("Content-Type")
	if !strings.Contains(contentTypeHeaderValue, "application/json") {
		http.Error(w, "unknown content-type", http.StatusNotImplemented)
		return
	}

	reader := json.NewDecoder(r.Body)
	reader.Decode(&agent)

	a, err := h.statusUC.GetUser(ctx, agent.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	buf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buf)
	err = jsonEncoder.Encode(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/repository"
	"github.com/IgorAleksandroff/agent-status/internal/usecase"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "login"
)

func (h *handler) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)

		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		login, err := h.auth.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		r.Header.Set(userCtx, login)
		next.ServeHTTP(w, r)
	})
}

func (h *handler) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
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

	newUser := entity.Agent{}
	reader := json.NewDecoder(r.Body)
	if err := reader.Decode(&newUser); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.auth.CreateUser(ctx, newUser); err != nil {
		log.Println(err.Error())
		if errors.Is(err, repository.ErrUserRegister) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := h.auth.GenerateToken(ctx, newUser.Login, newUser.Password)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(authorizationHeader, fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
}

func (h *handler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contentTypeHeaderValue := r.Header.Get("Content-Type")
	if !strings.Contains(contentTypeHeaderValue, "application/json") {
		http.Error(w, "unknown content-type", http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	user := entity.Agent{}
	reader := json.NewDecoder(r.Body)
	if err := reader.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.auth.GenerateToken(ctx, user.Login, user.Password)
	if err != nil {
		errLogin := errors.Is(err, usecase.ErrUserLogin)
		if errLogin {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(authorizationHeader, fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"payments/models"
	"payments/service"

	"github.com/zeebo/errs"
)

var authErr = errs.Class("authorization handlers")

type AuthorizationHandlers struct {
	service service.Authorization
}

func NewAuthorizationHandlers(s service.Authorization) *AuthorizationHandlers {
	return &AuthorizationHandlers{service: s}
}

func (h *AuthorizationHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		NewErrorResponse(w, authErr.Wrap(err), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		NewErrorResponse(w, authErr.Wrap(err), http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateUser(r.Context(), input)
	if err != nil {
		NewErrorResponse(w, authErr.Wrap(err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
	log.Printf("Handler: You are signed in as %v\n", id)

}

func (h *AuthorizationHandlers) Login(w http.ResponseWriter, r *http.Request) {
}

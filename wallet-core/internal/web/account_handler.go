package web

import (
	"encoding/json"
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/create_account"
	"log"
	"net/http"
)

type WebAccountHandler struct {
	CreateAccountUseCase create_account.CreateAccountUseCase
}

func NewWebAccountHandler(createAccountUseCase create_account.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{CreateAccountUseCase: createAccountUseCase}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto create_account.CreateAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	output, err := h.CreateAccountUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error executing use case: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error encoding response body: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

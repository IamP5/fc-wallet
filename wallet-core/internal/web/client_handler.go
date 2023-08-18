package web

import (
	"encoding/json"
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/create_client"
	"log"
	"net/http"
)

type WebClientHandler struct {
	CreateClientUseCase create_client.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase create_client.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto create_client.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	output, err := h.CreateClientUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error executing use case: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error encoding response body: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

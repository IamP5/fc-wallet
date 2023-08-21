package web

import (
	"encoding/json"
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/create_transaction"
	"log"
	"net/http"
)

type WebTransactionHandler struct {
	CreateTranscationUseCase create_transaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(createTranscationUseCase create_transaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{CreateTranscationUseCase: createTranscationUseCase}
}

func (h *WebTransactionHandler) CreateTranscation(w http.ResponseWriter, r *http.Request) {
	var dto create_transaction.CreateTransactionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	ctx := r.Context()
	output, err := h.CreateTranscationUseCase.Execute(ctx, dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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

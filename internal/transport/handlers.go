package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kiryshaaaa/infotecs-app/internal/dto"
)

func (h *APIHandlers) GetLastHandler(w http.ResponseWriter, r *http.Request) {
	countParam := r.URL.Query().Get("count")

	count, err := strconv.Atoi(countParam)
	if err != nil || count <= 0 {
		http.Error(w, "Invalid count parameter", http.StatusBadRequest)
		return
	}

	transactions, err := h.transactionService.GetLastN(count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *APIHandlers) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	log.Printf("Extracted address: %s", address)

	response, err := h.walletService.GetBalance(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *APIHandlers) SendHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.SendRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.FromWallet == "" || req.ToWallet == "" || req.Amount <= 0 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	err := h.transactionService.Send(req.FromWallet, req.ToWallet, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Transaction successful"}`))
}

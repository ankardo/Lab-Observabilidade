package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ankardo/Lab-Observabilidade/service-a/internal/usecases"
)

type ZipcodeHandler struct {
	Usecase *usecases.SendZipcode
}

func NewZipcodeHandler(usecase *usecases.SendZipcode) *ZipcodeHandler {
	return &ZipcodeHandler{Usecase: usecase}
}

func (h *ZipcodeHandler) SendZipcode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		CEP string `json:"cep"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.Usecase.Execute(req.CEP)
	if err != nil {
		if err.Error() == "invalid zipcode" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

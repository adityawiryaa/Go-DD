package application

import (
	"encoding/json"
	"errors"
	"go-driven-design/domain/entity"
	"go-driven-design/domain/usecase"
	"go-driven-design/pkg/common"
	"net/http"
)

type GreetHandler struct {
	greetingUseCase usecase.GreetingUseCase
}

func NewGreetHandler(greetingUseCase usecase.GreetingUseCase) *GreetHandler {
	return &GreetHandler{
		greetingUseCase: greetingUseCase,
	}

}

func (h *GreetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		err := errors.New("method not allowed")
		json.NewEncoder(w).Encode(common.ErrorResponse(
			http.StatusBadRequest,
			err.Error(),
			err,
		))
		return
	}
	var req *entity.GreetingDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := errors.New("invalid request")
		json.NewEncoder(w).Encode(common.ErrorResponse(
			http.StatusBadRequest,
			err.Error(),
			err,
		))
		return
	}
	result, err := h.greetingUseCase.Greet(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.ErrorResponse(
			0,
			err.Error(),
			err,
		))
		return
	}
	json.NewEncoder(w).Encode(common.SuccessResponse(
		"Sukses",
		0,
		result,
	))
}

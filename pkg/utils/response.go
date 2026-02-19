package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v3"
)

type DataResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteResponse(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"data": payload,
	})
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, err error) {
	httplog.SetAttrs(r.Context(),
		slog.Any("error", map[string]any{
			"message": err.Error(),
			"type":    fmt.Sprintf("%T", err),
			"status":  status,
		}),
	)

	WriteErrorResponse(w, http.StatusInternalServerError, err)
}

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"error": err.Error(),
	})
}

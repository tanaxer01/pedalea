package http

import (
	"net/http"

	"github.com/tanaxer01/pedalea/internal/core/bike"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type BikeHandler struct {
	service *bike.Service
}

func NewBikeHandler(service *bike.Service) *BikeHandler {
	return &BikeHandler{service: service}
}

func (h *BikeHandler) ListAvailableBikes(w http.ResponseWriter, r *http.Request) {
	bikes, err := h.service.ListAvailableBikes()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, bikes)
}

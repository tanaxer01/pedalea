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

// ListAvailableBikes godoc
// @Summary List available bikes
// @Description Returns all bikes that are currently available
// @Tags bikes
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.DataResponse{data=[]pedalea.Bike} "List of bikes in data"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /bikes/available [get]
func (h *BikeHandler) ListAvailableBikes(w http.ResponseWriter, r *http.Request) {
	bikes, err := h.service.ListAvailableBikes()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, bikes)
}

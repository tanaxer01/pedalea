package http

import (
	"net/http"
	"strconv"

	"github.com/tanaxer01/pedalea/internal/core/rental"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type RentalHandler struct {
	service *rental.Service
}

func NewRentalHandler(service *rental.Service) *RentalHandler {
	return &RentalHandler{service: service}
}

func (h *RentalHandler) StartRental(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("UserID").(string)
	if !ok {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, pedalea.ErrInvalidJwtSubject)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.StartRental](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.StartRental(intId, input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *RentalHandler) EndRental(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("UserID").(string)
	if !ok {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, pedalea.ErrInvalidJwtSubject)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.EndRental](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.EndRental(intId, input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *RentalHandler) ListUserRentals(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("UserID").(string)
	if !ok {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, pedalea.ErrInvalidJwtSubject)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	rentals, err := h.service.ListRentals(intId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, rentals)
}

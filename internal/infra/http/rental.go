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

// StartRental godoc
// @Summary Start a rental
// @Description Starts a bike rental for the authenticated user
// @Tags rentals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body pedalea.StartRental true "Rental start payload"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /rentals/start [get]
func (h *RentalHandler) StartRental(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("UserID").(string)
	if !ok {
		utils.WriteError(w, r, http.StatusInternalServerError, pedalea.ErrInvalidJwtSubject)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.StartRental](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.StartRental(intId, input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

// EndRental godoc
// @Summary End a rental
// @Description Ends the current rental for the authenticated user
// @Tags rentals
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body pedalea.EndRental true "Rental end payload"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /rentals/end [get]
func (h *RentalHandler) EndRental(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("UserID").(string)
	if !ok {
		utils.WriteError(w, r, http.StatusInternalServerError, pedalea.ErrInvalidJwtSubject)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.EndRental](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.EndRental(intId, input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

// ListUserRentals godoc
// @Summary List user rentals
// @Description Returns rental history for the authenticated user
// @Tags rentals
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.DataResponse{data=[]pedalea.RentalData} "Rental history in data"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /rentals/history [get]
func (h *RentalHandler) ListUserRentals(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("UserID").(string)
	if !ok {
		utils.WriteError(w, r, http.StatusInternalServerError, pedalea.ErrInvalidJwtSubject)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	rentals, err := h.service.ListRentals(intId)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, rentals)
}

package http

import (
	"net/http"
	"strconv"

	"github.com/tanaxer01/pedalea/internal/core/admin"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type AdminHandler struct {
	service *admin.Service
}

func NewAdminHandler(service *admin.Service) *AdminHandler {
	return &AdminHandler{service: service}
}

// Bikes
func (h *AdminHandler) InsertBike(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.BikeData](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.InsertBike(input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *AdminHandler) UpdateBike(w http.ResponseWriter, r *http.Request) {
	BikeID := r.PathValue("bike_id")
	ID, err := strconv.Atoi(BikeID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.BikeData](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateBike(ID, input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *AdminHandler) ListBikes(w http.ResponseWriter, r *http.Request) {
	bikes, err := h.service.ListBikes()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, bikes)
}

// User
func (h *AdminHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	UserID := r.PathValue("user_id")
	ID, err := strconv.Atoi(UserID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.UserData](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateUser(ID, input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *AdminHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, users)
}

func (h *AdminHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	UserID := r.PathValue("user_id")
	ID, err := strconv.Atoi(UserID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.GetUserData(ID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, user)
}

// Rental
func (h *AdminHandler) UpdateRental(w http.ResponseWriter, r *http.Request) {
	RentalID := r.PathValue("rental_id")
	ID, err := strconv.Atoi(RentalID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.RentalData](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateRental(ID, input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *AdminHandler) ListRentals(w http.ResponseWriter, r *http.Request) {
	rentals, err := h.service.ListRentals()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, rentals)
}

func (h *AdminHandler) GetRental(w http.ResponseWriter, r *http.Request) {
	RentalID := r.PathValue("rental_id")
	ID, err := strconv.Atoi(RentalID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	rental, err := h.service.GetRentalData(ID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, rental)
}

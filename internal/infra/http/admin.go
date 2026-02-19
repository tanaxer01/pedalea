package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
// InsertBike godoc
// @Summary Create a bike
// @Description Inserts a new bike (admin only)
// @Tags admin-bikes
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param payload body pedalea.BikeData true "Bike data"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/bikes [post]
func (h *AdminHandler) InsertBike(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.BikeData](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.InsertBike(input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

// UpdateBike godoc
// @Summary Update a bike
// @Description Updates a bike by ID (admin only)
// @Tags admin-bikes
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param bike_id path int true "Bike ID"
// @Param payload body pedalea.BikeData true "Bike data"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/bikes/{bike_id} [patch]
func (h *AdminHandler) UpdateBike(w http.ResponseWriter, r *http.Request) {
	bikeID := chi.URLParam(r, "bike_id")
	ID, err := strconv.Atoi(bikeID)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.BikeData](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateBike(ID, input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

// ListBikes godoc
// @Summary List bikes
// @Description Lists all bikes (admin only)
// @Tags admin-bikes
// @Produce json
// @Security BasicAuth
// @Success 200 {object} utils.DataResponse{data=[]pedalea.Bike} "List of bikes in data"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/bikes [get]
func (h *AdminHandler) ListBikes(w http.ResponseWriter, r *http.Request) {
	bikes, err := h.service.ListBikes()
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, bikes)
}

// User
// UpdateUser godoc
// @Summary Update a user
// @Description Updates a user by ID (admin only)
// @Tags admin-users
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param user_id path int true "User ID"
// @Param payload body pedalea.UserData true "User data"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/users/{user_id} [patch]
func (h *AdminHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	ID, err := strconv.Atoi(userID)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.UserData](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateUser(ID, input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

// ListUsers godoc
// @Summary List users
// @Description Lists all users (admin only)
// @Tags admin-users
// @Produce json
// @Security BasicAuth
// @Success 200 {object} utils.DataResponse{data=[]pedalea.User} "List of users in data"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/users [get]
func (h *AdminHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, users)
}

// GetUser godoc
// @Summary Get user
// @Description Gets a user by ID (admin only)
// @Tags admin-users
// @Produce json
// @Security BasicAuth
// @Param user_id path int true "User ID"
// @Success 200 {object} utils.DataResponse{data=pedalea.User} "User data in data"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/users/{user_id} [get]
func (h *AdminHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	ID, err := strconv.Atoi(userID)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.GetUserData(ID)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, user)
}

// Rental
// UpdateRental godoc
// @Summary Update a rental
// @Description Updates a rental by ID (admin only)
// @Tags admin-rentals
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param rental_id path int true "Rental ID"
// @Param payload body pedalea.RentalData true "Rental data"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/rentals/{rental_id} [patch]
func (h *AdminHandler) UpdateRental(w http.ResponseWriter, r *http.Request) {
	rentalID := chi.URLParam(r, "rental_id")
	ID, err := strconv.Atoi(rentalID)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	input, err := utils.ValidateBody[pedalea.RentalData](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateRental(ID, input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

// ListRentals godoc
// @Summary List rentals
// @Description Lists all rentals (admin only)
// @Tags admin-rentals
// @Produce json
// @Security BasicAuth
// @Success 200 {object} utils.DataResponse{data=[]pedalea.Rental} "List of rentals in data"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/rentals [get]
func (h *AdminHandler) ListRentals(w http.ResponseWriter, r *http.Request) {
	rentals, err := h.service.ListRentals()
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, rentals)
}

// GetRental godoc
// @Summary Get rental
// @Description Gets a rental by ID (admin only)
// @Tags admin-rentals
// @Produce json
// @Security BasicAuth
// @Param rental_id path int true "Rental ID"
// @Success 200 {object} utils.DataResponse{data=pedalea.Rental} "Rental data in data"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /admin/rentals/{rental_id} [get]
func (h *AdminHandler) GetRental(w http.ResponseWriter, r *http.Request) {
	rentalID := chi.URLParam(r, "rental_id")
	ID, err := strconv.Atoi(rentalID)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	rental, err := h.service.GetRentalData(ID)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, rental)
}

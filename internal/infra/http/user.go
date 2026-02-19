package http

import (
	"net/http"
	"strconv"

	"github.com/tanaxer01/pedalea/internal/core/user"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(service *user.Service) *UserHandler {
	return &UserHandler{service: service}
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags user
// @Accept json
// @Produce json
// @Param payload body pedalea.InsertUser true "User registration payload"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /user/register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.InsertUser](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.InsertUser(input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

// Login godoc
// @Summary Login user
// @Description Authenticates user and returns a JWT
// @Tags user
// @Accept json
// @Produce json
// @Param payload body pedalea.LoginUser true "Login payload"
// @Success 200 {object} utils.DataResponse{data=string} "JWT token in data"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /user/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.LoginUser](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Login(input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, token)
}

// GetUserData godoc
// @Summary Get current user profile
// @Description Returns the authenticated user's profile
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.DataResponse{data=pedalea.UserData} "User data in data"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /user/profile [get]
func (h *UserHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
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

	userData, err := h.service.GetUserData(intId)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, userData)
}

// UpdateUser godoc
// @Summary Update current user profile
// @Description Updates the authenticated user's profile
// @Tags user
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body pedalea.UserData true "User profile data"
// @Success 200 {object} utils.DataResponse{data=string} "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad request"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /user/profile [patch]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	input, err := utils.ValidateBody[pedalea.UserData](r.Body)
	if err != nil {
		utils.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateUser(intId, input)
	if err != nil {
		utils.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, map[string]string{"status": "ok"})
}

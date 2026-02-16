package http

import (
	"encoding/json"
	"net/http"

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

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.InsertUser](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
	}

	err = h.service.InsertUser(input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.LoginUser](r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}

	token, err := h.service.Login(input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, token)
}

func (h *UserHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	// TODO: Add claim data to the request context
	id, ok := r.Context().Value("sub").(int)
	if !ok {
		// TODO
	}

	userData, err := h.service.GetUserData(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, userData)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("sub").(int)
	if !ok {
		// TODO
	}

	input, err := utils.ValidateBody[pedalea.UserData](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
	}

	err = h.service.UpdateUser(id, input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}
}

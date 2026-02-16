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

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.InsertUser](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.InsertUser(input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	input, err := utils.ValidateBody[pedalea.LoginUser](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Login(input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, token)
}

func (h *UserHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
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

	userData, err := h.service.GetUserData(intId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, userData)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	input, err := utils.ValidateBody[pedalea.UserData](r.Body)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateUser(intId, input)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}
}

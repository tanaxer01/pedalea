package http

import (
	"net/http"

	"github.com/tanaxer01/pedalea/internal/core/admin"
	"github.com/tanaxer01/pedalea/pkg/utils"
)

type AdminHandler struct {
	service *admin.Service
}

func NewAdminHandler(service *admin.Service) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *AdminHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
	}

	utils.WriteResponse(w, users)
}

func (h *AdminHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// user, err := h.service.GetUserData()
}

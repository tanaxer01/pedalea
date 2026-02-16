package http

import (
	"github.com/tanaxer01/pedalea/internal/core/rental"
)

type RentalHandler struct {
	service *rental.Service
}

func NewRentalHandler(service *rental.Service) *RentalHandler {
	return &RentalHandler{service: service}
}

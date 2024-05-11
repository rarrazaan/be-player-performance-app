package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	service "github.com/rarrazaan/be-player-performance-app/internal/services"
)

type Handler struct {
}

func NewHandler(
	service *service.Service,
	cfg config.Config,
	validator *validator.Validate,
) *Handler {
	return &Handler{}
}

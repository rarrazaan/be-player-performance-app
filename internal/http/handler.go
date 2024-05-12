package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	service "github.com/rarrazaan/be-player-performance-app/internal/services"
)

type Handler struct {
	AuthHandler                *authHandler
	IdentityPerformanceHandler *identityPerformanceHandler
}

func NewHandler(
	service *service.Service,
	cfg config.Config,
	validator *validator.Validate,
) *Handler {
	return &Handler{
		AuthHandler:                NewAuthHandler(service.AuthService, cfg),
		IdentityPerformanceHandler: NewIdentityPerformanceHandler(service.IdentityPerformanceService, validator),
	}
}

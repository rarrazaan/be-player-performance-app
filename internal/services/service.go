package service

import (
	"github.com/rarrazaan/be-player-performance-app/internal/repository"
	"github.com/rarrazaan/be-player-performance-app/internal/utils"
)

type Service struct {
}

func NewService(
	repository *repository.Repository,
	jwt utils.IJWT,
) *Service {
	return &Service{}
}

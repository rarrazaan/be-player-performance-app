package repository

import (
	"github.com/rarrazaan/be-player-performance-app/internal/dependency/log"
	"gorm.io/gorm"
)

type IMonoRepository interface {
}
type monoRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewMonoRepository(db *gorm.DB, logger *log.Logger) IMonoRepository {
	return &monoRepository{
		db:     db,
		logger: logger,
	}
}

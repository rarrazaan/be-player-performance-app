package repository

import (
	"context"

	"github.com/rarrazaan/be-player-performance-app/internal/dependency/log"
	"github.com/rarrazaan/be-player-performance-app/internal/model"
	"gorm.io/gorm"
)

type IMonoRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
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

func (r *monoRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		r.logger.Error("MonoRepository.FindUserByEmail: " + err.Error())
		return nil, err
	}
	return user, nil
}

func (r *monoRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

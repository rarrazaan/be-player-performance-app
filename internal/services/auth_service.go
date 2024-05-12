package service

import (
	"context"
	"strings"

	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	"github.com/rarrazaan/be-player-performance-app/internal/model"
	"github.com/rarrazaan/be-player-performance-app/internal/repository"
	"github.com/rarrazaan/be-player-performance-app/internal/utils"
)

type IAuthService interface {
	LoginWithGoogle(ctx context.Context, googleUser *dto.GoogleResponse) (*dto.LoginResponsePayload, error)
}

func NewAuthservice(mr repository.IMonoRepository, jwt utils.IJWT) IAuthService {
	return &authService{
		mr:  mr,
		jwt: jwt,
	}
}

type authService struct {
	mr  repository.IMonoRepository
	cfg config.Config
	jwt utils.IJWT
}

func (s *authService) LoginWithGoogle(ctx context.Context, googleUser *dto.GoogleResponse) (*dto.LoginResponsePayload, error) {
	email := strings.ToLower(googleUser.Email)
	user, err := s.mr.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		newUser := &model.User{
			Username: email,
			Email:    email,
		}

		if _, err = s.mr.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}
		user, err = s.mr.FindUserByEmail(ctx, email)
		if err != nil {
			return nil, err
		}
	}

	accessTokenSignPayload := utils.SignAccessTokenPayload{
		UserID:    user.ID,
		UserName:  user.Username,
		UserEmail: user.Email,
	}
	accessToken, err := s.jwt.GenerateAccessToken(accessTokenSignPayload)
	if err != nil {
		return nil, err
	}

	responsePayload := &dto.LoginResponsePayload{
		AccessToken: *accessToken,
	}

	return responsePayload, nil
}

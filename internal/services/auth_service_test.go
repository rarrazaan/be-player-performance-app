package service

import (
	"context"
	"errors"
	"testing"

	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	"github.com/rarrazaan/be-player-performance-app/internal/model"
	"github.com/rarrazaan/be-player-performance-app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockMonoRepo *mocks.IMonoRepository
	mockJWT      *mocks.IJWT
	u            IAuthService
)

func Setup() {
	mockMonoRepo = new(mocks.IMonoRepository)
	mockJWT = new(mocks.IJWT)
	u = NewAuthservice(mockMonoRepo, mockJWT)
}
func Test_authService_LoginWithGoogle(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return correct user when email does exist", func(t *testing.T) {
		Setup()
		email := "test@gmail.com"
		user := &model.User{
			ID:    "",
			Email: email,
		}

		token := "token"
		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(user, nil)
		mockJWT.On("GenerateAccessToken", mock.Anything).Return(&token, nil)
		req := &dto.GoogleResponse{
			ID:            "",
			Email:         email,
			VerifiedEmail: true,
			Picture:       "",
		}
		res, err := u.LoginWithGoogle(context.Background(), req)
		expectedRes := &dto.LoginResponsePayload{
			AccessToken: "token",
		}

		assert.Equal(res, expectedRes)
		assert.Nil(err)
	})

	t.Run("should return user when email does not exist program will create it", func(t *testing.T) {
		Setup()
		email := "test@gmail.com"
		user := &model.User{
			ID:       "",
			Username: email,
			Email:    email,
		}

		token := "token"
		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(nil, nil).Once()
		mockMonoRepo.On("CreateUser", mock.Anything, user).Return(user, nil)
		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(user, nil).Once()
		mockJWT.On("GenerateAccessToken", mock.Anything).Return(&token, nil)

		req := &dto.GoogleResponse{
			ID:            "",
			Email:         email,
			VerifiedEmail: true,
			Picture:       "",
		}
		res, err := u.LoginWithGoogle(context.Background(), req)
		expectedRes := &dto.LoginResponsePayload{
			AccessToken: "token",
		}

		assert.Equal(res, expectedRes)
		assert.Nil(err)
	})

	t.Run("should return error when error happen beacuse of repository", func(t *testing.T) {
		Setup()
		email := "test@gmail.com"

		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(nil, errors.New("error happen"))
		req := &dto.GoogleResponse{
			ID:            "",
			Email:         email,
			VerifiedEmail: true,
			Picture:       "",
		}
		res, err := u.LoginWithGoogle(context.Background(), req)

		assert.NotNil(err)
		assert.Nil(res)
	})

	t.Run("should return error when error happen beacuse of repository (CreateUser)", func(t *testing.T) {
		Setup()
		email := "test@gmail.com"
		user := &model.User{
			ID:       "",
			Username: email,
			Email:    email,
		}

		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(nil, nil).Once()
		mockMonoRepo.On("CreateUser", mock.Anything, user).Return(nil, errors.New("error happen"))

		req := &dto.GoogleResponse{
			ID:            "",
			Email:         email,
			VerifiedEmail: true,
			Picture:       "",
		}
		res, err := u.LoginWithGoogle(context.Background(), req)

		assert.NotNil(err)
		assert.Nil(res)
	})

	t.Run("should return error when error happen beacuse of repository (second FindUserByEmail)", func(t *testing.T) {
		Setup()
		email := "test@gmail.com"
		user := &model.User{
			ID:       "",
			Username: email,
			Email:    email,
		}

		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(nil, nil).Once()
		mockMonoRepo.On("CreateUser", mock.Anything, user).Return(user, nil)
		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(nil, errors.New("error happen")).Once()

		req := &dto.GoogleResponse{
			ID:            "",
			Email:         email,
			VerifiedEmail: true,
			Picture:       "",
		}
		res, err := u.LoginWithGoogle(context.Background(), req)

		assert.NotNil(err)
		assert.Nil(res)
	})

	t.Run("should return error when failed to generate JWT", func(t *testing.T) {
		Setup()
		email := "test@gmail.com"
		user := &model.User{
			ID:    "",
			Email: email,
		}

		mockMonoRepo.On("FindUserByEmail", mock.Anything, email).Return(user, nil)
		mockJWT.On("GenerateAccessToken", mock.Anything).Return(nil, errors.New("error happen"))
		req := &dto.GoogleResponse{
			ID:            "",
			Email:         email,
			VerifiedEmail: true,
			Picture:       "",
		}
		res, err := u.LoginWithGoogle(context.Background(), req)

		assert.NotNil(err)
		assert.Nil(res)
	})
}

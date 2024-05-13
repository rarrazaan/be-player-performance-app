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
	mockMonoRepoIP *mocks.IMonoRepository
	ipu            IIdentityPerformanceService
)

func SetupIP() {
	mockMonoRepoIP = new(mocks.IMonoRepository)
	ipu = NewIdentityPerformanceService(mockMonoRepoIP)
}
func Test_identityPerformanceService_CalculatePerformance(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return correct result", func(t *testing.T) {
		SetupIP()
		performance := &dto.PerformanceRequest{
			N: 4,
			M: 2,
			A: "8 9 3 2",
			B: "5 4 1 3",
		}

		res := ipu.CalculatePerformance(context.Background(), performance)
		expectedRes := &dto.PerformanceResponse{
			Result: 6,
		}
		assert.Equal(res, expectedRes)
	})

	t.Run("should return nil when there are non integer element in array A", func(t *testing.T) {
		SetupIP()
		performance := &dto.PerformanceRequest{
			N: 4,
			M: 2,
			A: "8 a 3 2",
			B: "5 4 1 3",
		}

		res := ipu.CalculatePerformance(context.Background(), performance)

		assert.Nil(res)
	})

	t.Run("should return nil when there are non integer element in array B", func(t *testing.T) {
		SetupIP()
		performance := &dto.PerformanceRequest{
			N: 4,
			M: 2,
			A: "8 2 3 2",
			B: "5 a 1 3",
		}

		res := ipu.CalculatePerformance(context.Background(), performance)

		assert.Nil(res)
	})
}

func Test_identityPerformanceService_Identity(t *testing.T) {
	assert := assert.New(t)
	t.Run("should return correct identity when firstname exist", func(t *testing.T) {
		SetupIP()
		user := &dto.IdentityRequest{
			Name: "test",
		}
		identity := []model.UserDetail{
			{
				FullName: "test",
			},
		}

		mockMonoRepoIP.On("FindUserByFirstName", mock.Anything, user.Name).Return(identity, nil)

		res, err := ipu.Identity(context.Background(), user)
		expectedRes := []dto.IdentityResponse{
			{
				FullName: "test",
			},
		}

		assert.Equal(res, expectedRes)
		assert.Nil(err)
	})

	t.Run("should return error when error happen beacuse of repository", func(t *testing.T) {
		SetupIP()
		user := &dto.IdentityRequest{
			Name: "test",
		}

		mockMonoRepoIP.On("FindUserByFirstName", mock.Anything, user.Name).Return(nil, errors.New("error happen"))

		res, err := ipu.Identity(context.Background(), user)

		assert.NotNil(err)
		assert.Nil(res)
	})
}

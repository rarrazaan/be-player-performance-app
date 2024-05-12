package service

import (
	"context"
	"sort"
	"strconv"
	"strings"

	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	"github.com/rarrazaan/be-player-performance-app/internal/repository"
)

type IIdentityPerformanceService interface {
	CalculatePerformance(ctx context.Context, performance *dto.PerformanceRequest) *dto.PerformanceResponse
	Identity(ctx context.Context, performance *dto.IdentityRequest) ([]dto.IdentityResponse, error)
}

func NewIdentityPerformanceService(mr repository.IMonoRepository) IIdentityPerformanceService {
	return &identityPerformanceService{
		mr: mr,
	}
}

type identityPerformanceService struct {
	mr repository.IMonoRepository
}

type helper struct {
	numA int
	numB int
}

func (s *identityPerformanceService) CalculatePerformance(ctx context.Context, performance *dto.PerformanceRequest) *dto.PerformanceResponse {
	a := strings.Split(performance.A, " ")
	b := strings.Split(performance.B, " ")

	newArr := make([]helper, 0, performance.N)

	for i := 0; i < performance.N; i++ {
		aInt, err := strconv.Atoi(a[i])
		if err != nil {
			return nil
		}

		bInt, err := strconv.Atoi(b[i])
		if err != nil {
			return nil
		}

		newArr = append(newArr, helper{
			numA: aInt,
			numB: bInt,
		})
	}
	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i].numA < newArr[j].numA
	})

	res := &dto.PerformanceResponse{
		Result: performance.M,
	}
	for _, elm := range newArr {
		if elm.numA <= res.Result {
			res.Result += elm.numB
			continue
		}
		break
	}
	return res
}

func (s *identityPerformanceService) Identity(ctx context.Context, performance *dto.IdentityRequest) ([]dto.IdentityResponse, error) {
	users, err := s.mr.FindUserByFirstName(ctx, performance.Name)
	if err != nil {
		return nil, err
	}
	res := make([]dto.IdentityResponse, 0)
	for _, user := range users {
		res = append(res, dto.IdentityResponse{
			FullName:    user.FullName,
			Age:         user.Age,
			Gender:      user.Gender,
			Address:     user.Address,
			PhoneNumber: user.PhoneNumber,
		})
	}
	return res, nil
}

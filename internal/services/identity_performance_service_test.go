package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	"github.com/rarrazaan/be-player-performance-app/internal/repository"
)

func Test_identityPerformanceService_CalculatePerformance(t *testing.T) {
	type fields struct {
		mr repository.IMonoRepository
	}
	type args struct {
		ctx         context.Context
		performance *dto.PerformanceRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *dto.PerformanceResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &identityPerformanceService{
				mr: tt.fields.mr,
			}
			if got := s.CalculatePerformance(tt.args.ctx, tt.args.performance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("identityPerformanceService.CalculatePerformance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identityPerformanceService_Identity(t *testing.T) {
	type fields struct {
		mr repository.IMonoRepository
	}
	type args struct {
		ctx         context.Context
		performance *dto.IdentityRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.IdentityResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &identityPerformanceService{
				mr: tt.fields.mr,
			}
			got, err := s.Identity(tt.args.ctx, tt.args.performance)
			if (err != nil) != tt.wantErr {
				t.Errorf("identityPerformanceService.Identity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("identityPerformanceService.Identity() = %v, want %v", got, tt.want)
			}
		})
	}
}

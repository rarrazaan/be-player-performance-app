package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"github.com/rarrazaan/be-player-performance-app/internal/dto"
	"github.com/rarrazaan/be-player-performance-app/internal/repository"
	"github.com/rarrazaan/be-player-performance-app/internal/utils"
)

func Test_authService_LoginWithGoogle(t *testing.T) {
	type fields struct {
		mr  repository.IMonoRepository
		cfg config.Config
		jwt utils.IJWT
	}
	type args struct {
		ctx        context.Context
		googleUser *dto.GoogleResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.LoginResponsePayload
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authService{
				mr:  tt.fields.mr,
				cfg: tt.fields.cfg,
				jwt: tt.fields.jwt,
			}
			got, err := s.LoginWithGoogle(tt.args.ctx, tt.args.googleUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.LoginWithGoogle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.LoginWithGoogle() = %v, want %v", got, tt.want)
			}
		})
	}
}

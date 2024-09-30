package signup

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/service/users"
	usersStorage "github.com/egor-zakharov/goph-keeper/internal/storage/users"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()
	testIn := models.User{
		Login:    "1",
		Password: "1",
	}
	svrOut := models.User{
		UserID:   "1",
		Login:    "1",
		Password: "1",
	}
	err := errors.New("some err")
	type fields struct {
		usersService func(ctrl *gomock.Controller) users.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.SignUpRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				usersService: func(ctrl *gomock.Controller) users.Service {
					mock := users.NewMockService(ctrl)
					mock.EXPECT().Register(ctx, testIn).Return(&svrOut, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.SignUpRequest{
					Login:    "1",
					Password: "1",
				},
			},
			wantErr: false,
		},
		{
			name: "error from service",
			fields: fields{
				usersService: func(ctrl *gomock.Controller) users.Service {
					mock := users.NewMockService(ctrl)
					mock.EXPECT().Register(ctx, testIn).Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.SignUpRequest{
					Login:    "1",
					Password: "1",
				},
			},
			wantErr: true,
		},
		{
			name: "error from service duplicate login",
			fields: fields{
				usersService: func(ctrl *gomock.Controller) users.Service {
					mock := users.NewMockService(ctrl)
					mock.EXPECT().Register(ctx, testIn).Return(nil, usersStorage.ErrConflict).Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.SignUpRequest{
					Login:    "1",
					Password: "1",
				},
			},
			wantErr: true,
		},
		{
			name: "error validation login",
			fields: fields{
				usersService: func(ctrl *gomock.Controller) users.Service {
					mock := users.NewMockService(ctrl)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.SignUpRequest{
					Login:    "",
					Password: "1",
				},
			},
			wantErr: true,
		},
		{
			name: "error validation password",
			fields: fields{
				usersService: func(ctrl *gomock.Controller) users.Service {
					mock := users.NewMockService(ctrl)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.SignUpRequest{
					Login:    "1",
					Password: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := New(tt.fields.usersService(ctrl))
			_, err := h.Handle(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

package users

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/users"
	"github.com/egor-zakharov/goph-keeper/internal/utils"

	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_service_Login(t *testing.T) {
	ctx := context.Background()

	userIn := models.User{
		Login:    "login",
		Password: "password",
	}

	password, _ := utils.GetHashPassword(userIn.Password)

	userOut := models.User{
		UserID:   "1",
		Login:    "login",
		Password: password,
	}

	err := errors.New("some error")

	type fields struct {
		storage func(ctrl *gomock.Controller) users.Storage
	}
	type args struct {
		ctx    context.Context
		userIn models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) users.Storage {

					mock := users.NewMockStorage(ctrl)
					mock.EXPECT().Login(gomock.Any(), userIn.Login).Return(&userOut, nil)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				userIn: userIn,
			},
			want:    &userOut,
			wantErr: false,
		},
		{
			name: "error from storage",
			fields: fields{
				storage: func(ctrl *gomock.Controller) users.Storage {

					mock := users.NewMockStorage(ctrl)
					mock.EXPECT().Login(gomock.Any(), userIn.Login).Return(nil, err)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				userIn: userIn,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := &service{
				storage: tt.fields.storage(ctrl),
			}
			got, err := s.Login(tt.args.ctx, tt.args.userIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Register(t *testing.T) {
	ctx := context.Background()

	userIn := models.User{
		Login:    "login",
		Password: "password",
	}

	password, _ := utils.GetHashPassword(userIn.Password)

	userOut := models.User{
		UserID:   "1",
		Login:    "login",
		Password: password,
	}

	err := errors.New("some error")

	type fields struct {
		storage func(ctrl *gomock.Controller) users.Storage
	}
	type args struct {
		ctx    context.Context
		userIn models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) users.Storage {
					mock := users.NewMockStorage(ctrl)
					mock.EXPECT().Register(gomock.Any(), gomock.Any()).Return(&userOut, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				userIn: userIn,
			},
			want:    &userOut,
			wantErr: false,
		},
		{
			name: "error from storage",
			fields: fields{
				storage: func(ctrl *gomock.Controller) users.Storage {
					mock := users.NewMockStorage(ctrl)
					mock.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				userIn: userIn,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := &service{
				storage: tt.fields.storage(ctrl),
			}
			got, err := s.Register(tt.args.ctx, tt.args.userIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package updateauthdata

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/service/authdata"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.WithValue(context.Background(), auth.UserIDContextKey, "1")
	err := errors.New("some error")

	type fields struct {
		service      func(ctrl *gomock.Controller) authdata.Service
		notification func(ctrl *gomock.Controller) notification.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.UpdateAuthDataRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateAuthDataResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				service: func(ctrl *gomock.Controller) authdata.Service {
					mock := authdata.NewMockService(ctrl)
					mock.EXPECT().Update(ctx, models.AuthData{
						ID:       "1",
						Meta:     "2",
						Login:    "2",
						Password: "2",
					}, "1").Return(&models.AuthData{
						ID:       "1",
						Meta:     "2",
						Login:    "2",
						Password: "2",
					}, nil).Times(1)
					return mock
				},
				notification: func(ctrl *gomock.Controller) notification.Service {
					mock := notification.NewMockService(ctrl)
					mock.EXPECT().Send(ctx, notification.ProductAuth, notification.ActionUpdate, "1").Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.UpdateAuthDataRequest{
					Data: &pb.UpdateAuthDataRequest_Data{
						Id:       "1",
						Meta:     "2",
						Login:    "2",
						Password: "2",
					},
				},
			},
			want: &pb.UpdateAuthDataResponse{
				Result: true,
			},
			wantErr: false,
		},
		{
			name: "error from service",
			fields: fields{
				service: func(ctrl *gomock.Controller) authdata.Service {
					mock := authdata.NewMockService(ctrl)
					mock.EXPECT().Update(ctx, models.AuthData{
						ID:       "1",
						Meta:     "2",
						Login:    "2",
						Password: "2",
					}, "1").Return(nil, err).Times(1)
					return mock
				},
				notification: func(ctrl *gomock.Controller) notification.Service {
					mock := notification.NewMockService(ctrl)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.UpdateAuthDataRequest{
					Data: &pb.UpdateAuthDataRequest_Data{
						Id:       "1",
						Meta:     "2",
						Login:    "2",
						Password: "2",
					},
				},
			},
			want: &pb.UpdateAuthDataResponse{
				Result: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := New(tt.fields.service(ctrl), tt.fields.notification(ctrl))
			got, err := h.Handle(tt.args.ctx, tt.args.in)
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

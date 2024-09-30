package deletetextdata

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	"github.com/egor-zakharov/goph-keeper/internal/service/textdata"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.WithValue(context.Background(), auth.UserIDContextKey, "1")
	err := errors.New("some error")
	type fields struct {
		service      func(ctrl *gomock.Controller) textdata.Service
		notification func(ctrl *gomock.Controller) notification.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.DeleteTextDataRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteTextDataResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				service: func(ctrl *gomock.Controller) textdata.Service {
					mock := textdata.NewMockService(ctrl)
					mock.EXPECT().Delete(ctx, "1", "1").Return(nil).Times(1)
					return mock
				},
				notification: func(ctrl *gomock.Controller) notification.Service {
					mock := notification.NewMockService(ctrl)
					mock.EXPECT().Send(ctx, notification.ProductText, notification.ActionDelete, "1").Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.DeleteTextDataRequest{
					Id: "1",
				},
			},
			want: &pb.DeleteTextDataResponse{
				Result: true,
			},
			wantErr: false,
		},
		{
			name: "error from service",
			fields: fields{
				service: func(ctrl *gomock.Controller) textdata.Service {
					mock := textdata.NewMockService(ctrl)
					mock.EXPECT().Delete(ctx, "1", "1").Return(err).Times(1)
					return mock
				},
				notification: func(ctrl *gomock.Controller) notification.Service {
					mock := notification.NewMockService(ctrl)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.DeleteTextDataRequest{
					Id: "1",
				},
			},
			want: &pb.DeleteTextDataResponse{
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

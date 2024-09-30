package gettextdata

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/models"
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
		service func(ctrl *gomock.Controller) textdata.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.GetTextDataRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetTextDataResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				service: func(ctrl *gomock.Controller) textdata.Service {
					mock := textdata.NewMockService(ctrl)
					mock.EXPECT().Read(ctx, "1").Return(&[]models.TextData{{
						ID:   "1",
						Meta: "1",
						Text: "1",
					}}, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in:  nil,
			},
			want: &pb.GetTextDataResponse{
				Data: []*pb.GetTextDataResponse_Data{{
					Id:   "1",
					Meta: "1",
					Text: "1",
				}},
			},
			wantErr: false,
		},
		{
			name: "error from service",
			fields: fields{
				service: func(ctrl *gomock.Controller) textdata.Service {
					mock := textdata.NewMockService(ctrl)
					mock.EXPECT().Read(ctx, "1").Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in:  nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := New(tt.fields.service(ctrl))
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

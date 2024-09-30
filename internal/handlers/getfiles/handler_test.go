package getfiles

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/service/files"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.WithValue(context.Background(), auth.UserIDContextKey, "1")
	err := errors.New("some error")

	type fields struct {
		service func(ctrl *gomock.Controller) files.Service
	}
	type args struct {
		ctx context.Context
		in1 *pb.GetFilesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetFilesResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				service: func(ctrl *gomock.Controller) files.Service {
					mock := files.NewMockService(ctrl)
					mock.EXPECT().Read(ctx, "1").Return(&[]models.FileData{{
						ID:   "1",
						Name: "1",
						Meta: "1",
					}}, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in1: nil,
			},
			want: &pb.GetFilesResponse{
				Files: []*pb.GetFilesResponse_File{{
					Id:   "1",
					Name: "1",
					Meta: "1",
				}},
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				service: func(ctrl *gomock.Controller) files.Service {
					mock := files.NewMockService(ctrl)
					mock.EXPECT().Read(ctx, "1").Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in1: nil,
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
			got, err := h.Handle(tt.args.ctx, tt.args.in1)
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

package createcard

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/service/cards"
	"github.com/egor-zakharov/goph-keeper/internal/service/notification"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.WithValue(context.Background(), auth.UserIDContextKey, "1")

	type fields struct {
		cards        func(ctrl *gomock.Controller) cards.Service
		notification func(ctrl *gomock.Controller) notification.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateCardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateCardResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				cards: func(ctrl *gomock.Controller) cards.Service {
					mock := cards.NewMockService(ctrl)
					mock.EXPECT().Create(ctx, models.Card{
						Number:         "5576835778031510",
						ExpirationDate: "11/11",
						HolderName:     "1",
						CVV:            "1",
					}, "1").Return(&models.Card{
						ID:             "1",
						Number:         "1",
						ExpirationDate: "1",
						HolderName:     "1",
						CVV:            "1",
					}, nil).Times(1)
					return mock
				},
				notification: func(ctrl *gomock.Controller) notification.Service {
					mock := notification.NewMockService(ctrl)
					mock.EXPECT().Send(ctx, notification.ProductCard, notification.ActionCreate, "1")
					return mock
				},
			},
			args: args{
				ctx: ctx,
				in: &pb.CreateCardRequest{
					Card: &pb.CreateCardRequest_Card{
						Number:         "5576835778031510",
						ExpirationDate: "11/11",
						HolderName:     "1",
						Cvv:            "1",
					},
				},
			},
			want: &pb.CreateCardResponse{
				CardID: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := New(tt.fields.cards(ctrl), tt.fields.notification(ctrl))
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

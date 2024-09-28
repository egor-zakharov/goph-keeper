package cards

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/cards"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_service_Create(t *testing.T) {
	ctx := context.Background()
	testIn := models.Card{
		Number:         "5576835778031510",
		ExpirationDate: "11/11",
		HolderName:     "1",
		CVV:            "111",
	}
	testOut := models.Card{
		ID:             "5576835778031510",
		Number:         "",
		ExpirationDate: "11/11",
		HolderName:     "1",
		CVV:            "111",
	}
	userID := "1"
	err := errors.New("some err")
	type fields struct {
		storage func(ctrl *gomock.Controller) cards.Storage
	}
	type args struct {
		ctx    context.Context
		card   models.Card
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Card
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Create(gomock.Any(), testIn, userID).Return(&testOut, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				card:   testIn,
				userID: userID,
			},
			want:    &testOut,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Create(gomock.Any(), testIn, userID).Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				card:   testIn,
				userID: userID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := New(tt.fields.storage(ctrl))
			got, err := s.Create(tt.args.ctx, tt.args.card, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Delete(t *testing.T) {
	ctx := context.Background()
	userID := "1"
	cardID := "1"
	err := errors.New("some err")
	type fields struct {
		storage func(ctrl *gomock.Controller) cards.Storage
	}
	type args struct {
		ctx    context.Context
		cardID string
		userID string
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
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Delete(gomock.Any(), cardID, userID).Return(nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				cardID: cardID,
				userID: userID,
			},
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Delete(gomock.Any(), cardID, userID).Return(err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				cardID: cardID,
				userID: userID,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := New(tt.fields.storage(ctrl))
			if err := s.Delete(tt.args.ctx, tt.args.cardID, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Read(t *testing.T) {
	ctx := context.Background()
	testOut := []models.Card{
		{
			ID:             "5576835778031510",
			Number:         "",
			ExpirationDate: "11/11",
			HolderName:     "1",
			CVV:            "111",
		},
	}
	userID := "1"
	err := errors.New("some err")
	type fields struct {
		storage func(ctrl *gomock.Controller) cards.Storage
	}
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.Card
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Read(gomock.Any(), userID).Return(&testOut, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want:    &testOut,
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Read(gomock.Any(), userID).Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := New(tt.fields.storage(ctrl))
			got, err := s.Read(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Update(t *testing.T) {
	ctx := context.Background()
	testIn := models.Card{
		Number:         "5576835778031510",
		ExpirationDate: "11/11",
		HolderName:     "1",
		CVV:            "111",
	}
	userID := "1"
	err := errors.New("some err")
	type fields struct {
		storage func(ctrl *gomock.Controller) cards.Storage
	}
	type args struct {
		ctx    context.Context
		card   models.Card
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Card
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Update(gomock.Any(), testIn, userID).Return(&testIn, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				card:   testIn,
				userID: userID,
			},
			want:    &testIn,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				storage: func(ctrl *gomock.Controller) cards.Storage {
					mock := cards.NewMockStorage(ctrl)
					mock.EXPECT().Update(gomock.Any(), testIn, userID).Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				card:   testIn,
				userID: userID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			s := New(tt.fields.storage(ctrl))
			got, err := s.Update(tt.args.ctx, tt.args.card, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

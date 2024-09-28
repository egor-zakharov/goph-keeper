package authdata

import (
	"context"
	"errors"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	"github.com/egor-zakharov/goph-keeper/internal/storage/authdata"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_service_Create(t *testing.T) {
	ctx := context.Background()
	err := errors.New("some err")
	userID := "1"
	testIn := models.AuthData{
		Meta:     "1",
		Login:    "1",
		Password: "1",
	}
	testOut := models.AuthData{
		ID:       "1",
		Meta:     "1",
		Login:    "1",
		Password: "1",
	}
	type fields struct {
		storage func(ctrl *gomock.Controller) authdata.Storage
	}
	type args struct {
		ctx      context.Context
		authData models.AuthData
		userID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.AuthData
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
					mock.EXPECT().Create(gomock.Any(), testIn, userID).Return(&testOut, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:      ctx,
				authData: testIn,
				userID:   userID,
			},
			want:    &testOut,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
					mock.EXPECT().Create(gomock.Any(), testIn, userID).Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:      ctx,
				authData: testIn,
				userID:   userID,
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
			got, err := s.Create(tt.args.ctx, tt.args.authData, tt.args.userID)
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
	err := errors.New("some err")
	userID := "1"
	authID := "1"
	type fields struct {
		storage func(ctrl *gomock.Controller) authdata.Storage
	}
	type args struct {
		ctx    context.Context
		id     string
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
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
					mock.EXPECT().Delete(gomock.Any(), authID, userID).Return(nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				id:     authID,
				userID: userID,
			},
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
					mock.EXPECT().Delete(gomock.Any(), authID, userID).Return(err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:    ctx,
				id:     authID,
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
			if err := s.Delete(tt.args.ctx, tt.args.id, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Read(t *testing.T) {
	ctx := context.Background()
	err := errors.New("some err")
	userID := "1"
	testOut := []models.AuthData{
		{
			ID:       "1",
			Meta:     "1",
			Login:    "1",
			Password: "1",
		},
	}

	type fields struct {
		storage func(ctrl *gomock.Controller) authdata.Storage
	}
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.AuthData
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
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
			name: "error",
			fields: fields{
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
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
	err := errors.New("some err")
	userID := "1"
	testIn := models.AuthData{
		ID:       "1",
		Meta:     "1",
		Login:    "1",
		Password: "1",
	}
	type fields struct {
		storage func(ctrl *gomock.Controller) authdata.Storage
	}
	type args struct {
		ctx      context.Context
		authData models.AuthData
		userID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.AuthData
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
					mock.EXPECT().Update(gomock.Any(), testIn, userID).Return(&testIn, nil).Times(1)
					return mock
				},
			},
			args: args{
				ctx:      ctx,
				authData: testIn,
				userID:   userID,
			},
			want:    &testIn,
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				storage: func(ctrl *gomock.Controller) authdata.Storage {
					mock := authdata.NewMockStorage(ctrl)
					mock.EXPECT().Update(gomock.Any(), testIn, userID).Return(nil, err).Times(1)
					return mock
				},
			},
			args: args{
				ctx:      ctx,
				authData: testIn,
				userID:   userID,
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
			got, err := s.Update(tt.args.ctx, tt.args.authData, tt.args.userID)
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

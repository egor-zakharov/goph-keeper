package client

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/models"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"google.golang.org/grpc"
)

type Client interface {
	Connect() error
	SignUp(login string, password string) error
	SignIn(login string, password string) error
	CreateAuthData(ctx context.Context, meta string, login string, password string) (*models.AuthData, error)
	CreateCard(ctx context.Context, number string, expirationDate string, holderName string, CVV string) (*models.Card, error)
	CreateTextData(ctx context.Context, meta string, data string) (*models.TextData, error)
	DeleteAuthData(ctx context.Context, id string) error
	DeleteCard(ctx context.Context, id string) error
	DeleteFile(ctx context.Context, id string) error
	DeleteTextData(ctx context.Context, id string) error
	DownloadFile(ctx context.Context, id string, fileName string) error
	GetAuthData(ctx context.Context) (*[]models.AuthData, error)
	GetCards(ctx context.Context) (*[]models.Card, error)
	GetFiles(ctx context.Context) (*[]models.FileData, error)
	GetTextData(ctx context.Context) (*[]models.TextData, error)
	SubscribeToChanges(ctx context.Context) (grpc.ServerStreamingClient[pb.SubscribeToChangesResponse], error)
	UpdateAuthData(ctx context.Context, id string, meta string, login string, password string) error
	UpdateCard(ctx context.Context, id string, number string, expirationDate string, holderName string, CVV string) error
	UpdateTextData(ctx context.Context, id string, meta string, data string) error
	UploadFile(ctx context.Context, filePath string, meta string) (*models.FileData, error)
}

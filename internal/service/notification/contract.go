package notification

import (
	"context"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
)

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=notification

type Service interface {
	Add(ctx context.Context, stream pb.GophKeeperServer_SubscribeToChangesServer)
	Send(ctx context.Context, product string, action string, id string)
}

package notification

import (
	"context"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
)

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=notification

const (
	ProductCard = "card"
	ProductFile = "file"
	ProductText = "text"
	ProductAuth = "auth"
)

const (
	ActionCreate = "create"
	ActionDelete = "delete"
	ActionUpdate = "update"
)

type Service interface {
	Add(ctx context.Context, stream pb.GophKeeper_SubscribeToChangesServer)
	Send(ctx context.Context, product string, action string, id string)
}

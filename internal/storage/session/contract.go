package session

import pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"

//go:generate mockgen -source=contract.go -destination=contract_mock.go -package=session

type Storage interface {
	Add(userID string, sessionID string, stream pb.GophKeeper_SubscribeToChangesServer)
	Get(userID string, sessionID string) map[string]pb.GophKeeper_SubscribeToChangesServer
}

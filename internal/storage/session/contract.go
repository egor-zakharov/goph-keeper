package session

import pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"

type Storage interface {
	Add(userID string, sessionID string, stream pb.GophKeeperServer_SubscribeToChangesServer)
	Get(userID string, sessionID string) map[string]pb.GophKeeperServer_SubscribeToChangesServer
}

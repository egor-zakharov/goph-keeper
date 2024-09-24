package notification

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"github.com/egor-zakharov/goph-keeper/internal/storage/session"
)

type service struct {
	session session.Storage
}

func New(session session.Storage) Service {
	return &service{session: session}
}

func (s service) Add(ctx context.Context, stream pb.GophKeeper_SubscribeToChangesServer) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	sessionID, _ := ctx.Value(auth.SessionIdContextKey).(string)
	s.session.Add(userID, sessionID, stream)
}

func (s service) Send(ctx context.Context, product string, action string, id string) {
	userID := ctx.Value(auth.UserIdContextKey).(string)
	sessionID, _ := ctx.Value(auth.SessionIdContextKey).(string)
	get := s.session.Get(userID, sessionID)
	for session, client := range get {
		if session != sessionID {
			_ = client.Send(&pb.SubscribeToChangesResponse{
				Product: product,
				Action:  action,
				Id:      id,
			})
		}
	}
}

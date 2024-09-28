package notification

import (
	"context"
	"github.com/egor-zakharov/goph-keeper/internal/auth"
	"github.com/egor-zakharov/goph-keeper/internal/storage/session"
	pb "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
)

type service struct {
	session session.Storage
}

func New(session session.Storage) Service {
	return &service{session: session}
}

func (s service) Add(ctx context.Context, stream pb.GophKeeper_SubscribeToChangesServer) {
	userID := ctx.Value(auth.UserIDContextKey).(string)
	sessionID, _ := ctx.Value(auth.SessionIDContextKey).(string)
	s.session.Add(userID, sessionID, stream)
}

func (s service) Send(ctx context.Context, product string, action string, id string) {
	userID := ctx.Value(auth.UserIDContextKey).(string)
	sessionID, _ := ctx.Value(auth.SessionIDContextKey).(string)
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

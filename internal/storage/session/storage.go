package session

import (
	pb "github.com/egor-zakharov/goph-keeper/internal/proto/gophkeeper"
	"sync"
)

type storage struct {
	syncClients map[string]map[string]pb.GophKeeper_SubscribeToChangesServer
	mu          sync.RWMutex
}

func New() Storage {
	return &storage{
		syncClients: make(map[string]map[string]pb.GophKeeper_SubscribeToChangesServer),
	}
}

func (s *storage) Add(userID string, sessionID string, stream pb.GophKeeper_SubscribeToChangesServer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.syncClients[userID]; !exists {
		s.syncClients[userID] = make(map[string]pb.GophKeeper_SubscribeToChangesServer)
	}
	s.syncClients[userID][sessionID] = stream
}

func (s *storage) Get(userID string, sessionID string) map[string]pb.GophKeeper_SubscribeToChangesServer {
	s.mu.RLock()
	defer s.mu.RUnlock()

	response := make(map[string]pb.GophKeeper_SubscribeToChangesServer)

	for session, client := range s.syncClients[userID] {
		if session != sessionID {
			response[session] = client
		}
	}
	return response
}

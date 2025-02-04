package terminal

import (
	"context"
	"sync"

	uuid "github.com/google/uuid"
)

type TerminalManager struct {
	sessions map[string]*TerminalSession
	mu       sync.Mutex
}

func NewTerminalManager() *TerminalManager {
	return &TerminalManager{
		sessions: make(map[string]*TerminalSession),
	}
}

func (tm *TerminalManager) CreateSession(ctx context.Context) (*TerminalSession, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	id := uuid.New().String()
	session, err := NewTerminalSession(id, ctx)
	if err != nil {
		return nil, err
	}

	tm.sessions[id] = session

	return session, nil
}

func (tm *TerminalManager) GetSession(id string) *TerminalSession {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	session, exists := tm.sessions[id]
	if !exists {
		return nil
	}

	return session
}

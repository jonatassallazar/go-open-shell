package main

import (
	"context"
	"fmt"

	"go-open-shell/src/terminal"
)

// App struct
type App struct {
	ctx     context.Context
	manager *terminal.TerminalManager
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.manager = terminal.NewTerminalManager()
}

func (a *App) CreateSession() (terminal.TerminalSession, error) {
	session, err := a.manager.CreateSession(a.ctx)
	if err != nil {
		fmt.Printf("error: %s", err)
		return *session, err
	}

	return *session, nil
}

func (a *App) SendMessageInput(id string) terminal.TerminalSession {
	session := a.manager.GetSession(id)
	// session.Input.In = i.In
	session.SendMessageInput()

	return *session
}

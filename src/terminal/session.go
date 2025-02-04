package terminal

import (
	"context"
	"fmt"
	"os"
	"time"

	wails "github.com/wailsapp/wails/v2/pkg/runtime"
)

type InputEvent struct {
	In   string    `json:"in"`
	Out  string    `json:"out"`
	Err  error     `json:"err,omitempty"`
	Dir  string    `json:"dir"`
	Time time.Time `json:"time"`
}

type TerminalSession struct {
	ID          string          `json:"id"`
	Input       InputEvent      `json:"input"`
	History     []InputEvent    `json:"history"`
	Dir         string          `json:"dir"`
	PreviousDir string          `json:"previousDir,omitempty"`
	Context     context.Context `json:"ctx,omitempty"`
}

func (t *TerminalSession) Write(p []byte) (n int, err error) {
	t.Input.Out = t.Input.Out + string(p)

	fmt.Println(string(p))
	wails.EventsEmit(t.Context, "input-response", &t)

	return len(p), nil
}

func (t *TerminalSession) WriteError(err error) {
	t.Input.Err = err
}

func (t *TerminalSession) AppendHistory() {
	t.History = append(t.History, InputEvent{
		In:   t.Input.In,
		Out:  t.Input.Out,
		Err:  t.Input.Err,
		Dir:  t.PreviousDir,
		Time: time.Now(),
	})

	// t.Input = InputEvent{}
}

func NewTerminalSession(id string, ctx context.Context) (*TerminalSession, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	return &TerminalSession{
		ID:          id,
		Input:       InputEvent{},
		History:     []InputEvent{},
		Dir:         homedir,
		PreviousDir: homedir,
		Context:     ctx,
	}, nil
}

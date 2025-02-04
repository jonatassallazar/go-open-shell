package terminal

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	wails "github.com/wailsapp/wails/v2/pkg/runtime"
)

type EventData struct {
	Input InputEvent
}

func (ts *TerminalSession) SendMessageInput() {
	// create channel for keep func waiting for commands
	close := make(chan error)

	// set the current directory to the session directory
	err := os.Chdir(ts.Dir)
	if err != nil {
		ts.WriteError(err)
		return
	}

	var cmd *exec.Cmd

	// check if the OS is windows and use powershell
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "chcp 65001 > NUL && powershell ")
	} else {
		// cmd = exec.Command("sh", "-c", input)
	}

	// set the command directory, stdout, and stderr
	cmd.Dir = ts.Dir

	// create io pipes for streaming content
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("error creating stdin pipe: %s\n", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("error creating stdout pipe: %s\n", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("error creating stderr pipe: %s\n", err)
	}

	// start cmd
	if err := cmd.Start(); err != nil {
		fmt.Printf("error starting command: %s\n", err)
	}

	// copy stdout and stderr output to terminal session
	go io.Copy(ts, stdout)
	go io.Copy(ts, stderr)

	// Launch a goroutine to wait for the command to complete and send the response to the channel
	go func() {
		close <- cmd.Wait()
		fmt.Println("created routine to wait")
	}()

	// go routine that will wait for the shell commands
	go func() {
		wails.EventsOn(ts.Context, "send-input", func(optionalData ...interface{}) {
			fmt.Println("send-input", optionalData)
			if len(optionalData) > 0 {
				defer ts.AppendHistory()
				if i, ok := optionalData[0].(string); ok {
					fmt.Printf("write input: %s\n", i)
					io.WriteString(stdin, fmt.Sprintf("%s\n", i))
				}
			}
		})
	}()

	// fire init event so all listeners know that shell has opened
	wails.EventsEmit(ts.Context, "session-initiated")

	// Wait for the response from the channel before continuing
	if err := <-close; err != nil {
		fmt.Printf("error waiting for command: %s\n", err)
	}

}

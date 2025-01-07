package terminal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

func OpenTerminal() {
	reader := bufio.NewReader(os.Stdin)
	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
	}()

	fmt.Println("Welcome to Go Terminal! Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading the input.")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		if strings.HasPrefix(input, "cd ") {
			dir := strings.TrimSpace(strings.TrimPrefix(input, "cd "))
			err := os.Chdir(dir)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				currentDir, _ = os.Getwd()
			}
			continue
		}

		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell", "-Command", input)
		} else {
			cmd = exec.Command("sh", "-c", input)
		}
		cmd.Dir = currentDir
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var builtinCommands = []string{"echo", "exit", "type", "pwd"}

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Println("Error reading command")

		}

		command = strings.TrimRight(command, "\r\n")
		command = strings.TrimSpace(command)
		command, args := split(command)

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			echoCommand(args)
		case "type":
			typeCommand(args)
		case "pwd":
			pwdCommand()
		default:
			cmd := exec.Command(command, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				if errors.Is(err, exec.ErrNotFound) {
					fmt.Printf("%s: command not found\n", command)
				} else {
					log.Println("command failed")
				}
			}

		}
	}
}

func split(command string) (string, []string) {
	tokens := strings.Fields(command)
	if len(tokens) == 1 {
		return tokens[0], []string{}
	}

	return tokens[0], tokens[1:]

}

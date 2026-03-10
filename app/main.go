package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

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
			echo(args)
		default:
			notFound(command)
		}

		fmt.Printf("\n")
	}
}

func notFound(command string) {
	fmt.Printf("%s: command not found", command)
}

func split(command string) (string, []string) {
	tokens := strings.Fields(command)
	if len(tokens) == 1 {
		return tokens[0], []string{}
	}

	return tokens[0], tokens[1:]

}

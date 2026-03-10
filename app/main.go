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

		switch command {
		case "exit":
			exit()
		}

		notFound(command)
		fmt.Printf("\n")
	}
}

func notFound(command string) {
	fmt.Printf("%s: command not found", command)
}

func exit() {
	os.Exit(0)
}

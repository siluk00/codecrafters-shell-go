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

	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	command = strings.TrimRight(command, "\r\n")
	if err != nil {
		log.Println("Error reading command")

	}
	notFound(command)

}

func notFound(command string) {
	fmt.Printf("%s: command not found", command)
}

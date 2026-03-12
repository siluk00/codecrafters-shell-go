package main

import (
	"fmt"
	"log"
	"os"
)

func pwdCommand() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Couldn't execute command")
		os.Exit(1)
	}
	fmt.Printf("%s\n", dir)
}

package main

import (
	"fmt"
	"strings"
)

func echoCommand(args []string) {
	phrase := strings.Join(args, " ")
	fmt.Printf("%s\n", phrase)
}

package main

import (
	"fmt"
	"strings"
)

func echoCommand(args []string) {
	fmt.Printf("%s\n", strings.Join(args, " "))
}

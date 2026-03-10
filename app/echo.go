package main

import (
	"fmt"
	"strings"
)

func echo(args []string) {
	phrase := strings.Join(args, " ")
	fmt.Printf("%s", phrase)
}

package main

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"
)

func typeCommand(commands string) {
	args := strings.Fields(commands)
	for _, arg := range args {
		if slices.Contains(builtinCommands, arg) {
			fmt.Printf("%s is a shell builtin\n", arg)
		} else {
			if path, err := exec.LookPath(arg); err != nil {
				fmt.Printf("%s: not found\n", arg)
			} else {
				fmt.Printf("%s is %s\n", arg, path)
			}
		}
	}
}

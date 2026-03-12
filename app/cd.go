package main

import (
	"fmt"
	"os"
)

func (c *Cmd) cdCommand(args []string) {
	if len(args) > 1 {
		fmt.Println("cd: too many arguments\n")
	}
	if len(args) == 0 {
		return
	}
	if args[0][0] == '/' {
		fileInfo, err := os.Stat(args[0])

		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("cd: %s: No such file or directory/n", args[0])
			}
			return
		}

		if fileInfo.IsDir() {
			c.curDir = args[0]
			//fmt.Printf("%s\n", args[0])
		}
	} else {
		fmt.Printf("cd: %s: No such file or directory/n", args[0])

	}
}

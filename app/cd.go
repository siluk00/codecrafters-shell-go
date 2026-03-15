package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (c *Cmd) cdCommand(path string) {
	args := strings.Fields(path)

	if len(args) > 1 {
		fmt.Println("cd: too many arguments")
	}
	if len(args) == 0 {
		return
	}

	switch args[0][0] {
	case '/':
		path = args[0]
	case '.':
		path = filepath.Join(c.curDir, args[0])
	case '~':
		if args[0] == "~" {
			var err error
			path, err = os.UserHomeDir()
			if err != nil {
				return
			}
		}
	default:
		fmt.Printf("cd: %s: No such file or directory\n", args[0])
		return
	}

	fileInfo, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("cd: %s: No such file or directory\n", args[0])
		}
		return
	}

	if fileInfo.IsDir() {
		c.curDir = path
		return
		//fmt.Printf("%s\n", args[0])
	}

	fmt.Printf("cd: %s: No such file or directory\n", args[0])
}

package main

import (
	"fmt"
)

func (c *Cmd) pwdCommand() {
	fmt.Printf("%s\n", c.curDir)
}

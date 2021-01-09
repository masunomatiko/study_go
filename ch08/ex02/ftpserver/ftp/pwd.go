package ftp

import (
	"fmt"
)

func (c *Conn) pwd() {
	fmt.Println(c.workDir)
	c.respond(status200)
}

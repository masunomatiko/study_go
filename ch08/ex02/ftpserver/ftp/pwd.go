package ftp

import (
	"fmt"
	"log"
)

func (c *Conn) pwd() {

	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	fmt.Println(c.workDir)
	_, err = fmt.Fprint(dataConn, c.workDir, "\n")
	if err != nil {
		log.Print(err)
		c.respond(status426)
	}
	c.respond(status200)
}

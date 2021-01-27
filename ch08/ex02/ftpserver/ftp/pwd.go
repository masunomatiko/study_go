package ftp

import (
	"fmt"
	"log"
	"path/filepath"
)

func (c *Conn) pwd() {

	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	absPath := filepath.Join(c.rootDir, c.workDir)
	_, err = fmt.Fprint(dataConn, absPath, "\n")
	if err != nil {
		log.Print(err)
		c.respond(status426)
	}
	c.respond(status200)
}

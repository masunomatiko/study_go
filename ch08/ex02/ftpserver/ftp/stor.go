package ftp

import (
	"io"
	"os"
)

func (c *Conn) stor(args []string) {
	if len(args) != 1 {
		c.respond(status501)
	}
	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
	}
	fname := args[0]

	file, err := os.Create(fname)
	defer dataConn.Close()
	_, err = io.Copy(dataConn, file)
	if err != nil {
		c.respond(status426)
	}

	c.respond(status200)
}

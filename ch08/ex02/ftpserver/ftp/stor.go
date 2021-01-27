package ftp

import (
	"io"
	"os"
	"path/filepath"
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

	file, err := os.Create(filepath.Join(c.rootDir, c.workDir, fname))
	if err != nil {
		c.respond(status550)
	}
	defer dataConn.Close()
	_, err = io.Copy(dataConn, file)
	if err != nil {
		c.respond("450 File unavailable.")
	}

	c.respond(status200)
}

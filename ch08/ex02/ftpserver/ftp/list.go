package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func (c *Conn) list(args []string) {
	var target string
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := ioutil.ReadDir(target)
	if err != nil {
		log.Print(err)
		c.respond(status550)
		return
	}
	c.respond("150 File status okay; about to open data connection.")

	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
		return
	}
	defer dataConn.Close()

	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), "\n")
		if err != nil {
			log.Print(err)
			c.respond(status426)
		}
	}
	c.respond(status200)
}

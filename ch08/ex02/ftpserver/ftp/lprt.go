package ftp

import "log"

func (c *Conn) lprt(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}
	longIP, err := longIPFromHostPort(args[0])
	if err != nil {
		log.Print(err)
		c.respond(status425)
		return
	}
	c.longIP = longIP
	c.respond(status200)
}

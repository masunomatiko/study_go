package ftp

import (
	"fmt"
	"log"
)

const (
	status200 = "200 Command okay."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status550 = "550 Requested action not taken. File unavailable."
)

func (c *Conn) respond(s string) {
	_, err := fmt.Fprint(c.conn, s, "\n")
	if err != nil {
		log.Print(err)
	}
}

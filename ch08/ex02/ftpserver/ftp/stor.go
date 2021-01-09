package ftp

func (c *Conn) stor(args []string) {
	dataConn, err := c.dataConnect()
	if err != nil {
		c.respond(status425)
	}
	defer dataConn.Close()

	c.respond(status200)
}

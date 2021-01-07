package ftp

import "net"

// Connection to the FTP server
type Conn struct {
	conn     net.Conn
	longIP *longIP
	rootDir  string
	workDir  string
}

// NewConn returns a new FTP connection
func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}

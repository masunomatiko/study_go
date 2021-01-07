package ftp

import (
	"fmt"
	"net"
)

func (c *Conn) dataConnect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.longIP.toAddress())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type longIP struct {
	protocol                                                              int
	addrLength                                                            int
	h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11, h12, h13, h14, h15, h16 int // host
	portLength                                                            int
	p1, p2                                                                int // port
}

func longIPFromHostPort(hostPort string) (*longIP, error) {
	var lip longIP
	// PORTじゃなくてLPRTしか使えない
	_, err := fmt.Sscanf(hostPort, "%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
		&lip.protocol, &lip.addrLength, &lip.h1, &lip.h2, &lip.h3, &lip.h4, &lip.h5, &lip.h6, &lip.h7, &lip.h8, &lip.h9, &lip.h10, &lip.h11, &lip.h12, &lip.h13, &lip.h14, &lip.h15, &lip.h16, &lip.portLength, &lip.p1, &lip.p2)
	if err != nil {
		return nil, err
	}
	return &lip, nil
}

func (lip *longIP) toAddress() string {
	if lip == nil {
		return ""
	}
	// 実際のポート番号　＝　nnn × 256 + mmm
	port := lip.p1<<8 + lip.p2
	// IPv6がわからなくて挫折した。あとでやる。
	// return fmt.Sprintf("%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d:%d", lip.h1, lip.h2, lip.h3, lip.h4, lip.h5, lip.h6, lip.h7, lip.h8, lip.h9, lip.h10, lip.h11, lip.h12, lip.h13, lip.h14, lip.h15, lip.h16, port)
	return fmt.Sprintf("[0:0:0:0:0:0:0:1]:%d", port)

}

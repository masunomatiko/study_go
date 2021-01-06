package ftp

import (
	"fmt"
	"net"
)

func (c *Conn) dataConnect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.dataPort.toAddress())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type dataPort struct {
	protocol                                                              int
	addrLength                                                            int
	h1, h2, h3, h4, h5, h6, h7, h8, h9, h10, h11, h12, h13, h14, h15, h16 int // host
	portLength                                                            int
	p1, p2                                                                int // port
}

func dataPortFromHostPort(hostPort string) (*dataPort, error) {
	var dp dataPort
	// PORTじゃなくてLPRTしか使えない
	_, err := fmt.Sscanf(hostPort, "%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
		&dp.protocol, &dp.addrLength, &dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.h5, &dp.h6, &dp.h7, &dp.h8, &dp.h9, &dp.h10, &dp.h11, &dp.h12, &dp.h13, &dp.h14, &dp.h15, &dp.h16, &dp.portLength, &dp.p1, &dp.p2)
	if err != nil {
		return nil, err
	}
	return &dp, nil
}

func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}
	// 実際のポート番号　＝　nnn × 256 + mmm
	port := d.p1<<8 + d.p2
	return fmt.Sprintf("%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, d.h5, d.h6, d.h7, d.h8, d.h9, d.h10, d.h11, d.h12, d.h13, d.h14, d.h15, d.h16, port)
}

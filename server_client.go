package tcp

import "net"

type ServerClient struct {
	Id   int
	conn *net.TCPConn
}

package tcp

import (
	"net"
)

type OnNewServerMessage func(message []byte)

type TcpClient struct {
	parameters *ClientParameters
	conn       *net.TCPConn
}

func newTcpClient(parameters *ClientParameters) TcpClient {
	return TcpClient{
		parameters: parameters,
	}
}

func (tcp *TcpClient) Connect() error {
	serverAddr := tcp.parameters.Ip + ":" + tcp.parameters.Port
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)

	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		return err
	}

	tcp.conn = conn
	go tcp.bindInput()
	return nil
}

func (tcp *TcpClient) bindInput() {
	for {
		bytes := make([]byte, tcp.parameters.MaxSizeBuffer)
		_, err := tcp.conn.Read(bytes)

		if err == nil {
			tcp.parameters.OnNewServerMessage(bytes)
		}
	}
}

func (tcp *TcpClient) SendMessage(message []byte) {
	tcp.conn.Write([]byte(message))
}

func (tcp *TcpClient) Close() {
	tcp.conn.Close()
}

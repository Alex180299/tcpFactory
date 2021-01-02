package tcp

import (
	"bufio"
	"net"
)

type tcpClient struct {
	parameters *ClientParameters
	conn       *net.TCPConn
}

func newTcpClient(parameters *ClientParameters) Tcp {
	return &tcpClient{
		parameters: parameters,
	}
}

func (tcp *tcpClient) Connect() error {
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
	go tcp.bindInputChannel()
	return nil
}

func (tcp *tcpClient) bindInputChannel() {
	serverReader := bufio.NewReader(tcp.conn)
	for {
		serverResponse, err := serverReader.ReadString(tcp.parameters.Delimiter)

		if err == nil {
			tcp.parameters.InputChannel <- serverResponse
		}
	}
}

func (tcp *tcpClient) Close() {
	tcp.conn.Close()
}

package tcp

import (
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
	go tcp.bindOutputChannel()
	return nil
}

func (tcp *tcpClient) bindInputChannel() {
	for {
		bytes := make([]byte, tcp.parameters.MaxSizeBuffer)
		_, err := tcp.conn.Read(bytes)

		if err == nil {
			tcp.parameters.InputChannel <- string(bytes)
		}
	}
}

func (tcp *tcpClient) bindOutputChannel() {
	for {
		outputMessage := <-tcp.parameters.OutputChannel
		tcp.conn.Write([]byte(outputMessage))
	}
}

func (tcp *tcpClient) Close() {
	tcp.conn.Close()
}

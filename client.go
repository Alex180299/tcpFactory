package tcp

import (
	"fmt"
	"net"
	"time"
)

type OnNewServerMessage func(message []byte)

type TcpClient struct {
	parameters *ClientParameters
	tcpAddr    *net.TCPAddr
	conn       *net.TCPConn
}

func newTcpClient(parameters *ClientParameters) TcpClient {
	return TcpClient{
		parameters: parameters,
	}
}

func (tcp *TcpClient) Connect() error {
	serverAddr := tcp.parameters.Ip + ":" + tcp.parameters.Port
	errResolveAddr := tcp.resolveAddress(serverAddr)

	if errResolveAddr != nil {
		return errResolveAddr
	}

	errResolveTcp := tcp.resolveDialTCP()

	if errResolveTcp != nil {
		return errResolveTcp
	}

	go tcp.bindInput()
	return nil
}

func (tcp *TcpClient) resolveAddress(serverAddr string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)

	if err != nil {
		reconnectionTime := tcp.parameters.ConnectionParameters.ReconnectionTime
		if reconnectionTime > 0 {
			t, _ := time.ParseDuration(fmt.Sprintf("%dms", reconnectionTime))
			time.Sleep(t)
			fmt.Println("Error resolving address: ", serverAddr, ", reconnecting")
			return tcp.resolveAddress(serverAddr)
		} else {
			return err
		}
	} else {
		tcp.tcpAddr = tcpAddr
		return nil
	}
}

func (tcp *TcpClient) resolveDialTCP() error {
	conn, err := net.DialTCP("tcp", nil, tcp.tcpAddr)

	if err != nil {
		reconnectionTime := tcp.parameters.ConnectionParameters.ReconnectionTime
		if reconnectionTime > 0 {
			t, _ := time.ParseDuration(fmt.Sprintf("%dms", reconnectionTime))
			time.Sleep(t)
			fmt.Println("Error to connect client, reconnecting")
			return tcp.resolveDialTCP()
		} else {
			return err
		}
	} else {
		tcp.conn = conn
		return nil
	}
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

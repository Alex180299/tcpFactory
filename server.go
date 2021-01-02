package tcp

import (
	"net"
)

type OnNewClientListener func(client *ServerClient)

type tcpServer struct {
	parameters *ServerParameters
	listener   *net.TCPListener
}

func newTcpServer(parameters *ServerParameters) Tcp {
	return &tcpServer{
		parameters: parameters,
	}
}

func (tcp *tcpServer) Connect() error {
	serverAddr := tcp.parameters.Ip + ":" + tcp.parameters.Port

	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	tcp.listener = listener
	go tcp.acceptClient()
	return nil
}

func (tcp *tcpServer) acceptClient() {
	for {
		conn, err := tcp.listener.AcceptTCP()
		if err != nil {
			continue
		}

		go tcp.handleNewClient(conn)
	}
}

func (tcp *tcpServer) handleNewClient(conn *net.TCPConn) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)

	serverClient := &ServerClient{
		id:            "",
		InputChannel:  inputChannel,
		OutputChannel: outputChannel,
	}

	go tcp.parameters.OnNewClientListener(serverClient)
	go tcp.handleClientInputChannel(conn, serverClient)
	go tcp.handleClientOutputChannel(conn, serverClient)
}

func (tcp *tcpServer) handleClientInputChannel(conn *net.TCPConn, client *ServerClient) {
	for {
		bytes := make([]byte, tcp.parameters.MaxSizeBuffer)
		_, err := conn.Read(bytes)

		if err == nil {
			client.InputChannel <- string(bytes)
		}
	}
}

func (tcp *tcpServer) handleClientOutputChannel(conn *net.TCPConn, client *ServerClient) {
	for {
		outputMessage := <-client.OutputChannel
		conn.Write([]byte(outputMessage))
	}
}

func (tcp *tcpServer) Close() {
	tcp.listener.Close()
}

package tcp

import (
	"bufio"
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
		conn, err := tcp.listener.Accept()
		if err != nil {
			continue
		}

		go tcp.handleNewClient(conn)
	}
}

func (tcp *tcpServer) handleNewClient(conn net.Conn) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)

	serverClient := &ServerClient{
		id:            "",
		InputChannel:  inputChannel,
		OutputChannel: outputChannel,
	}

	go tcp.parameters.OnNewClientListener(serverClient)

	clientReader := bufio.NewReader(conn)
	for {
		bytes := make([]byte, tcp.parameters.MaxSizeBuffer)
		_, err := clientReader.Read(bytes)

		if err == nil {
			serverClient.InputChannel <- string(bytes)
		}
	}
}

func (tcp *tcpServer) Close() {
	tcp.listener.Close()
}

package tcp

import (
	"net"
)

type OnNewClient func(client *ServerClient)
type OnNewClientMessage func(message []byte, client ServerClient)

type TcpServer struct {
	parameters *ServerParameters
	listener   *net.TCPListener
	clients    []*ServerClient
}

func newTcpServer(parameters *ServerParameters) TcpServer {
	return TcpServer{
		parameters: parameters,
	}
}

func (tcp *TcpServer) Connect() error {
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

func (tcp *TcpServer) acceptClient() {
	for {
		conn, err := tcp.listener.AcceptTCP()
		if err != nil {
			continue
		}

		go tcp.handleNewClient(conn)
	}
}

func (tcp *TcpServer) handleNewClient(conn *net.TCPConn) {
	serverClient := &ServerClient{
		Id:   len(tcp.clients),
		conn: conn,
	}

	tcp.clients = append(tcp.clients, serverClient)
	go tcp.parameters.OnNewClient(serverClient)
	go tcp.handleClientInputChannel(conn, serverClient)
}

func (tcp *TcpServer) handleClientInputChannel(conn *net.TCPConn, client *ServerClient) {
	for {
		bytes := make([]byte, tcp.parameters.MaxSizeBuffer)
		_, err := conn.Read(bytes)

		if err == nil {
			tcp.parameters.OnNewClientMessage(bytes, *client)
		}
	}
}

func (tcp *TcpServer) SendMessageToAll(message []byte) {
	for _, client := range tcp.clients {
		client.conn.Write([]byte(message))
	}
}

func (tcp *TcpServer) SendMessageToClient(message []byte, clientId int) {
	client := tcp.clients[clientId]
	client.conn.Write([]byte(message))
}

func (tcp *TcpServer) Close() {
	for _, client := range tcp.clients {
		client.conn.Close()
	}
	tcp.listener.Close()
}

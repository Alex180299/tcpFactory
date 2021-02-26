package tcp

import (
	"fmt"
	"net"
	"time"
)

type OnNewClient func(client *ServerClient, tcp TcpServer)
type OnNewClientMessage func(message []byte, client ServerClient, tcp TcpServer)

type TcpServer struct {
	parameters *ServerParameters
	listener   *net.TCPListener
	clients    []*ServerClient
	tcpAddr    *net.TCPAddr
}

func newTcpServer(parameters *ServerParameters) TcpServer {
	return TcpServer{
		parameters: parameters,
	}
}

func (tcp *TcpServer) Connect() error {
	serverAddr := tcp.parameters.Ip + ":" + tcp.parameters.Port

	errResolveAddr := tcp.resolveAddress(serverAddr)

	if errResolveAddr != nil {
		return errResolveAddr
	}

	err := tcp.resolveListenTCP()
	if err != nil {
		return err
	}

	go tcp.acceptClient()
	return nil
}

func (tcp *TcpServer) resolveAddress(serverAddr string) error {
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

func (tcp *TcpServer) resolveListenTCP() error {
	listener, err := net.ListenTCP("tcp", tcp.tcpAddr)

	if err != nil {
		reconnectionTime := tcp.parameters.ConnectionParameters.ReconnectionTime
		if reconnectionTime > 0 {
			t, _ := time.ParseDuration(fmt.Sprintf("%dms", reconnectionTime))
			time.Sleep(t)
			fmt.Println("Error to connect server, reconnecting")
			return tcp.resolveListenTCP()
		} else {
			return err
		}
	} else {
		tcp.listener = listener
		return nil
	}
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
	go tcp.parameters.OnNewClient(serverClient, *tcp)
	go tcp.handleClientInputChannel(conn, serverClient)
}

func (tcp *TcpServer) handleClientInputChannel(conn *net.TCPConn, client *ServerClient) {
	for {
		bytes := make([]byte, tcp.parameters.MaxSizeBuffer)
		_, err := conn.Read(bytes)

		if err == nil {
			tcp.parameters.OnNewClientMessage(bytes, *client, *tcp)
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

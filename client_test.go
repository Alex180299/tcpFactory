package tcp

import (
	"testing"
)

func createTcpServer() *TcpServer {
	server := newTcpServer(&ServerParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Server connection",
			MaxSizeBuffer:    100,
		},
		OnNewClient: func(client *ServerClient) {
			println("Client was connected with id: ", client.Id)
		},
	})

	errServer := server.Connect()

	if errServer != nil {
		panic("Error to create server")
	}

	return &server
}

func TestNewTcpClient(t *testing.T) {
	server := createTcpServer()

	client := newTcpClient(&ClientParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Client connection",
			MaxSizeBuffer:    100,
		},
		OnNewServerMessage: func(message []byte) {
			println("Message received: ", string(message))
		},
	})

	errClient := client.Connect()

	if errClient != nil {
		t.Errorf("error to create tcp client")
	}

	client.SendMessage([]byte("Hello from tcp client"))
	client.Close()

	server.Close()
}

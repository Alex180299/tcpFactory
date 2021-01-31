package tcp

import (
	"fmt"
	"testing"
	"time"
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
			fmt.Println("Client was connected with id: ", client.Id)
		},
		OnNewClientMessage: func(message []byte, client ServerClient) {
			fmt.Println("Message received from client ", client.Id, ":", string(message))
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

	//Wait to server receive and print the new client before to send message
	time.Sleep(10 * time.Millisecond)
	client.SendMessage([]byte("Hello from tcp client"))
	//Wait to server receive and print the message
	time.Sleep(10 * time.Millisecond)
	client.Close()

	server.Close()
}

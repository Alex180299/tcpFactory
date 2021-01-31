package tcp

import (
	"testing"
)

func TestNewTcpServer(t *testing.T) {
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

	err := server.Connect()

	if err != nil {
		t.Errorf("error when server listening")
	}

	server.Close()
}

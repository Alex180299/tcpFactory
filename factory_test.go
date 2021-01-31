package tcp

import "testing"

func TestGetNewTcpServer(t *testing.T) {
	GetNewTcpServer(&ServerParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Server connection",
		},
		OnNewClient:        func(client *ServerClient) {},
		OnNewClientMessage: func(message []byte, client ServerClient) {},
	})
}

func TestGetNewTcpClient(t *testing.T) {
	GetNewTcpClient(&ClientParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Client connection",
		},
		OnNewServerMessage: func(message []byte) {},
	})
}

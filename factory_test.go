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
		OnNewClientListener: func(client *ServerClient) {},
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
		InputChannel:  nil,
		OutputChannel: nil,
	})
}

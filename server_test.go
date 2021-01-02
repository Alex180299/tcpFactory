package tcp

import "testing"

func TestNewTcpServer(t *testing.T) {
	server := newTcpServer(&ServerParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Server connection",
		},
		OnNewClientListener: nil,
	})

	err := server.Connect()

	if err != nil {
		t.Errorf("error when server listening")
	}

	server.Close()
}
package tcp

import (
	"testing"
)

func TestNewTcpClient(t *testing.T) {
	client := newTcpClient(&ClientParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Client connection",
		},
		InputChannel:  nil,
		OutputChannel: nil,
	})

	err := client.Connect()

	if err != nil {
		t.Errorf("error to create tcp client")
	}

	client.Close()
}

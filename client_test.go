package tcp

import (
	"fmt"
	"testing"
)

func TestNewTcpClient(t *testing.T) {
	inputChannel := make(chan string)

	client := newTcpClient(&ClientParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Client connection",
			Delimiter:        ' ',
		},
		InputChannel:  inputChannel,
		OutputChannel: nil,
	})

	err := client.Connect()

	fmt.Println(<-inputChannel)

	if err != nil {
		t.Errorf("error to create tcp client")
	}

	client.Close()
}

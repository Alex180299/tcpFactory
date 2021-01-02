package tcp

import (
	"fmt"
	"testing"
)

func TestNewTcpClient(t *testing.T) {
	inputChannel := make(chan string)
	outputChannel := make(chan string)

	client := newTcpClient(&ClientParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Client connection",
			MaxSizeBuffer:    100,
		},
		InputChannel:  inputChannel,
		OutputChannel: outputChannel,
	})

	err := client.Connect()

	outputChannel <- "Golang client connected"
	fmt.Println(<-inputChannel)

	if err != nil {
		t.Errorf("error to create tcp client")
	}

	client.Close()
}

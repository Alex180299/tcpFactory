package tcp

import (
	"fmt"
	"testing"
)

func TestNewTcpServer(t *testing.T) {
	waitChannel := make(chan string)
	server := newTcpServer(&ServerParameters{
		ConnectionParameters: ConnectionParameters{
			Ip:               "localhost",
			Port:             "5000",
			ReconnectionTime: 1000,
			Name:             "Server connection",
			MaxSizeBuffer:    100,
		},
		OnNewClientListener: func(client *ServerClient) {
			client.OutputChannel <- "Hello from golang server"
			str := <-client.InputChannel
			waitChannel <- str
		},
	})

	err := server.Connect()

	if err != nil {
		t.Errorf("error when server listening")
	}

	fmt.Println(<-waitChannel)

	server.Close()
}

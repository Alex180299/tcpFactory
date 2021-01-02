package tcp

type ServerClient struct {
	id            string
	InputChannel  chan string
	OutputChannel chan string
}

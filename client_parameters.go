package tcp

type ClientParameters struct {
	ConnectionParameters
	InputChannel  chan string
	OutputChannel chan string
}

package tcp

type ServerParameters struct {
	ConnectionParameters
	OnNewClient
	OnNewClientMessage
}

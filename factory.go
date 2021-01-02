package tcp

func GetNewTcpServer(parameters *ServerParameters) Tcp {
	return newTcpServer(parameters)
}

func GetNewTcpClient(parameters *ClientParameters) Tcp {
	return newTcpClient(parameters)
}

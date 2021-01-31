package tcp

func GetNewTcpServer(parameters *ServerParameters) TcpServer {
	return newTcpServer(parameters)
}

func GetNewTcpClient(parameters *ClientParameters) TcpClient {
	return newTcpClient(parameters)
}

package tcp

type Tcp interface {
	Connect() error
	Close()
}

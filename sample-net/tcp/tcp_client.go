package tcp_net

type TcpClientState int8

const (
	ClientStop     TcpClientState = 0
	ClientTryStart TcpClientState = 1
	ClientRunning  TcpClientState = 2
	ClientTryStop  TcpClientState = 3
)

type TcpClient struct {
	// socket addr
	IP   string
	Port int
	// Client control
	exitChan  chan struct{}
	StateChan chan TcpClientState
}

type ITcpClient interface {
	// Client control
	Start()
	Stop()
}

func NewTcpClient(IP string, Port int) *TcpClient {
	c := TcpClient{
		IP:        IP,
		Port:      Port,
		exitChan:  make(chan struct{}),
		StateChan: make(chan TcpClientState),
	}
	return &c
}

func (c *TcpClient) Start() {
}

func (c *TcpClient) Stop() {
	c.StateChan <- ClientTryStop
	c.exitChan <- struct{}{}
}

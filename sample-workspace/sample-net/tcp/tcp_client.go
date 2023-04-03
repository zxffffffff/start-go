package tcp_net

import (
	"bufio"
	"fmt"
	"net"

	pb "github.com/zxffffffff/start-go/sample-pb/pb_go"
	"google.golang.org/protobuf/proto"
)

type TcpClientState int8

type TcpClient struct {
	// socket addr
	IP   string
	Port int
	// Client control
	exitChan  chan struct{}
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
	}
	return &c
}

// dial
func (c *TcpClient) Start() {
	fmt.Println("TcpClient Start")

	go func() {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.IP, c.Port))
		if err != nil {
			fmt.Println("TcpClient Dial err: ", err)
			return
		}

		// connection
		go func() {
			go func() {
				req1 := pb.HeartbeatReq{
					ReqIndex: 123,
				}
				buf, err := proto.Marshal(&req1)
				if err != nil {
					return
				}
				_, err = conn.Write(buf)
				if err != nil {
					return
				}
			}()

			reader := bufio.NewReader(conn)
			var buf [128]byte
			n, err := reader.Read(buf[:])
			if err != nil {
				fmt.Println("TcpClient Read err: ", err)
			}
			recv := string(buf[:n])
			fmt.Println("TcpClient Read: ", recv)
            // ping-pong
            conn.Write([]byte(recv))
		}()
    
        select {
        case <-c.exitChan:
            err := conn.Close()
            if err != nil {
                fmt.Println("TcpClient Close err: ", err)
            }
        }
	}()
}

func (c *TcpClient) Stop() {
	fmt.Println("TcpClient Stop")
	c.exitChan <- struct{}{}
}

package tcp_net

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

type TcpServerState int8

type TcpServer struct {
	// socket addr
	IP   string
	Port int
	// server control
	exitChan  chan struct{}
}

type ITcpServer interface {
	// server control
	Start()
	Stop()
}

func NewTcpServer(IP string, Port int) *TcpServer {
	s := TcpServer{
		IP:        IP,
		Port:      Port,
		exitChan:  make(chan struct{}),
	}
	return &s
}

func (s *TcpServer) Start() {
	fmt.Println("TcpServer Start")

	// listen
	go func() {
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("TcpServer ResolveTCPAddr err: ", err)
			return
		}

		listener, err := net.ListenTCP("tcp", addr)
		if err != nil {
			fmt.Println("TcpServer ListenTCP err: ", err)
			return
		}

		fmt.Println("TcpServer running: ", fmt.Sprintf("%s:%d", s.IP, s.Port))

		// accept
		go func() {
			for {
				conn, err := listener.AcceptTCP()
				if err != nil {
					if errors.Is(err, net.ErrClosed) {
						fmt.Println("TcpServer close")
						return
					}
					fmt.Println("TcpServer AcceptTCP err: ", err)
					continue
				}

				// connection
				go func() {
					reader := bufio.NewReader(conn)
					var buf [128]byte
					n, err := reader.Read(buf[:])
					if err != nil {
						fmt.Println("TcpServer Read err: ", err)
					}
					recv := string(buf[:n])
					fmt.Println("TcpServer Read: ", recv)
                    // ping-pong
					conn.Write([]byte(recv))
				}()
			}
		}()

		select {
		case <-s.exitChan:
			err := listener.Close()
			if err != nil {
				fmt.Println("TcpServer Close err: ", err)
			}
		}
	}()
}

func (s *TcpServer) Stop() {
	fmt.Println("TcpServer Stop")
	s.exitChan <- struct{}{}
}

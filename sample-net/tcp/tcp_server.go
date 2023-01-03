package tcp_net

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

type TcpServerState int8

const (
	ServerStop     TcpServerState = 0
	ServerTryStart TcpServerState = 1
	ServerRunning  TcpServerState = 2
	ServerTryStop  TcpServerState = 3
)

type TcpServer struct {
	// socket addr
	IP   string
	Port int
	// server control
	exitChan  chan struct{}
	StateChan chan TcpServerState
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
		StateChan: make(chan TcpServerState),
	}
	return &s
}

func (s *TcpServer) Start() {
    fmt.Println("TcpServer Start")
	s.StateChan <- ServerTryStart

	// listen
	go func() {
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ResolveTCPAddr err: ", err)
			s.StateChan <- ServerStop
			return
		}

		listener, err := net.ListenTCP("tcp", addr)
		if err != nil {
			fmt.Println("ListenTCP err: ", err)
			s.StateChan <- ServerStop
			return
		}

		fmt.Println("TcpServer running: ", fmt.Sprintf("%s:%d", s.IP, s.Port))
		s.StateChan <- ServerRunning

		// accept
		go func() {
			for {
				conn, err := listener.AcceptTCP()
				if err != nil {
					if errors.Is(err, net.ErrClosed) {
						fmt.Println("TcpServer close")
						s.StateChan <- ServerStop
						return
					}
					fmt.Println("AcceptTCP err: ", err)
					continue
				}

				// connection
				go onConn(s, conn)
			}
		}()

		select {
		case <-s.exitChan:
			err := listener.Close()
			if err != nil {
				fmt.Println("Close err: ", err)
			}
			s.StateChan <- ServerStop
		}
	}()
}

func (s *TcpServer) Stop() {
    fmt.Println("TcpServer Stop")
	s.StateChan <- ServerTryStop
	s.exitChan <- struct{}{}
}

func onConn(s *TcpServer, conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	var buf [128]byte
	n, err := reader.Read(buf[:])
	if err != nil {
		fmt.Println("Read err: ", err)
	}
	recv := string(buf[:n])
	fmt.Println("Read: ", recv)

	conn.Write([]byte(recv))
}

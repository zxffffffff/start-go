package net

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	go_net "net"
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
	s.StateChan <- ServerTryStart

	// listen
	go func() {
		addr, err := go_net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ResolveTCPAddr err: ", err)
			s.StateChan <- ServerStop
			return
		}

		listener, err := go_net.ListenTCP("tcp", addr)
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

				// connection (ping pong)
				go func() {
                    reader := bufio.NewReader(conn)
                    var buf[128]byte
                    n, err := reader.Read(buf[:])
                    if err != nil [
                        fmt.Println("Read err: ", err)
                    ]
                    recv := string(buf[:n])
                    fmt.Println("Read: ", recv)

                    conn.Write([]byte(recv))
				}()
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
	s.StateChan <- ServerTryStop
	s.exitChan <- struct{}{}
}

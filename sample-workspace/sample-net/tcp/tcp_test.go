package tcp_net

import (
	"testing"
	"time"
)

func TestGetMsg(t *testing.T) {
	server := NewTcpServer("127.0.0.1", 12345)
	server.Start()

	client := NewTcpClient("127.0.0.1", 12345)
	client.Start()

	//if err != nil {
	//	t.Fatalf("TestGetMsg error %q %v", msg, err)
	//}
	time.Sleep(time.Second * 3)
}

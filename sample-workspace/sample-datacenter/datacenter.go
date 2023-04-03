package datacenter

import (
	"errors"
	"fmt"

	net "github.com/zxffffffff/start-go/sample-net"
	"rsc.io/quote"
)

func GetMsg(key string) (string, error) {
	if key == "" {
		return "", errors.New("empty key")
	}

	msg := quote.Go()
	return fmt.Sprintf("[%v] %v", key, msg), nil
}

func GetNetMsg() {
	msg, err := net.GetMsg("test2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
}

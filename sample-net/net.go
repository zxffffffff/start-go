package xnet

import (
	"errors"
	"fmt"
	"rsc.io/quote"
)

func GetMsg(key string) (string, error) {
	if key == "" {
		return "", errors.New("empty key")
	}

	msg := quote.Go()
	return fmt.Sprintf("[%v] %v", key, msg), nil
}

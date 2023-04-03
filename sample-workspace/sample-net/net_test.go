package xnet

import (
	"regexp"
	"testing"
)

func TestGetMsg(t *testing.T) {
	key := "test"
	want := regexp.MustCompile(`[` + key + `]`)
	msg, err := GetMsg(key)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf("TestGetMsg error %q %v", msg, err)
	}
}

func TestGetMsgEmpty(t *testing.T) {
	msg, err := GetMsg("")
	if msg != "" || err == nil {
		t.Fatalf("TestGetMsgEmpty error %q %v", msg, err)
	}
}

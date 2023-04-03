package main

import (
	"fmt"
	"log"

	datacenter "github.com/zxffffffff/start-go/sample-datacenter"
)

func printMsg(key string) {
	msg, err := datacenter.GetMsg(key)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msg)
	}
}

func main() {
	log.SetPrefix("[log]: ")
	log.SetFlags(0)

	printMsg("test")
	printMsg("")
	datacenter.GetNetMsg()
}

package nnet

import (
	"log"
	"net"
	"testing"
)

func TestAddr(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(lis.Addr().String())
}

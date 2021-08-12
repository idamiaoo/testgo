package main

import (
	"net"

	"testgo/stdlib/tcp/kernelbuf/proto"

	log "github.com/sirupsen/logrus"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Errorf("accept: %v", err)
			continue
		}
		log.Infof("accept new client:%s", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		log.Infof("read")
		p := &proto.Proto{}
		if err := p.ReadFrom(conn); err != nil {
			log.Error(err)
			return
		}
		log.Info(p.String())
	}
}

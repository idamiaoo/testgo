package main

import (
	"net"
	"time"

	"testgo/stdlib/tcp/kernelbuf/proto"

	log "github.com/sirupsen/logrus"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.213.183:9999")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("connect success")
	i := 1
	for {
		p := &proto.Proto{
			Seq:  i,
			Body: "1111111111111111111111111111111111111111111111111111111111",
		}
		if err := p.WriteTo(conn); err != nil {
			log.Error(err)
			return
		} else {
			log.Infof("write success, seq:%d", i)
			i++
		}
		<-time.After(5 * time.Second)
	}
}

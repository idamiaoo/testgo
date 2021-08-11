package main

import (
	"net"

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
	/*
		root := fmt.Sprintf("/Users/user/go/src/github.com/katakurin/testgo/stdlib/tcp/bytes/server/%d", time.Now().Unix())

		_ = os.MkdirAll(root, 0777)
		for {
			f := &proto.File{}
			if err := f.ReadFile(conn); err != nil {
				log.Error(err)
				return
			}

			if err := os.WriteFile(filepath.Join(root, filepath.Base(f.Name)), f.Data, 0666); err != nil {
				log.Error(err)
				return
			}
			log.Infof("%s recived", f.Name)
		}
	*/

}

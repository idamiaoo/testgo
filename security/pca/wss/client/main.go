package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/websocket"
)

var (
	url    = "wss://ws.yiheng.life:8787/ws"
	origin = "http://localhost/"
)

func main() {
	config, err := websocket.NewConfig(url, origin)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../keys/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)

	clientCert, err := tls.LoadX509KeyPair("../../keys/client.crt", "../../keys/client.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair err:", err)
		return
	}
	config.TlsConfig = &tls.Config{
		RootCAs:      certPool,
		ServerName:   "ws.yiheng.life",
		Certificates: []tls.Certificate{clientCert},
	}
	conn, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := conn.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = conn.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}

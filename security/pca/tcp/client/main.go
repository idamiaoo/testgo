package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	addr = "ws.yiheng.life:8087"
)

func main() {
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
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		ServerName:   "ws.yiheng.life",
		Certificates: []tls.Certificate{clientCert},
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := conn.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	log.Println("write success")
	var msg = make([]byte, 512)
	var n int
	if n, err = conn.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}

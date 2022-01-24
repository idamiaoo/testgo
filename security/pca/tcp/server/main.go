package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../keys/ca.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	certPool.AppendCertsFromPEM(ca)

	serverCert, err := tls.LoadX509KeyPair("../../keys/server.crt", "../../keys/server.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair err:", err)
		return
	}

	tlsConfig := &tls.Config{
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{serverCert},
	}

	listener, err := tls.Listen("tcp", "0.0.0.0:8087", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			var msg = make([]byte, 512)
			var n int
			n, err := conn.Read(msg)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(string(msg[:n]))

			_, err = conn.Write([]byte("hello"))
			if err != nil {
				log.Println(err)
				return
			}
			conn.Close()
			log.Println("conn disconnect")
		}()
	}
}

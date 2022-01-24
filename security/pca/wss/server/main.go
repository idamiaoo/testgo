package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func echoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../keys/ca.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	certPool.AppendCertsFromPEM(ca)
	mux := http.NewServeMux()
	mux.Handle("/ws", websocket.Handler(echoServer))

	server := &http.Server{
		Addr:    ":8787",
		Handler: mux,
		TLSConfig: &tls.Config{
			ClientCAs: certPool,
			// ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	if err := server.ListenAndServeTLS("../../keys/server.crt", "../../keys/server.key"); err != nil {
		log.Fatal(err)
	}
}

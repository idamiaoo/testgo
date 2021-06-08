package main

import (
	"crypto/tls"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/websocket"
)

func handwss(conn *websocket.Conn) {
	conn.PayloadType = websocket.TextFrame
	host, port, err := net.SplitHostPort(conn.Request().RemoteAddr)
	if err != nil {
		log.Error("cannot get remote address:", err)
		return
	}
	log.Infof("new ws connection from:%v port: %v", host, port)
}

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.cmkj.com"), //your domain here
		Cache:      autocert.DirCache("certs"),             //folder for storing certificates
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	//http.Handle("/wss", websocket.Handler(handwss))

	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	if err := server.ListenAndServeTLS("", ""); err != nil { //key and cert are comming from Let's Encrypt
		log.Error(err)
	}
}

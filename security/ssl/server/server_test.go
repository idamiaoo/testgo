package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestServerVerificationSelf(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	server := &http.Server{
		Addr:    ":8787",
		Handler: mux,
	}

	if err := server.ListenAndServeTLS("../keys/cert.pem", "../keys/server.key"); err != nil {
		log.Fatal(err)
	}
}

func TestServerVerification(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	server := &http.Server{
		Addr:    ":8787",
		Handler: mux,
	}

	if err := server.ListenAndServeTLS("../keys/ca/server.crt", "../keys/ca/server.key"); err != nil {
		log.Fatal(err)
	}
}

func TestClientVerification(t *testing.T) {
	pool := x509.NewCertPool()
	caCertPath := "../keys/ca/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	server := &http.Server{
		Addr:    ":8787",
		Handler: mux,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	if err := server.ListenAndServeTLS("../keys/ca/server.crt", "../keys/ca/server.key"); err != nil {
		log.Fatal(err)
	}
}

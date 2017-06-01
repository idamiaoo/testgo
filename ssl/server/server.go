package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	server := &http.Server{Addr: ":8787", Handler: mux}

	if err := server.ListenAndServeTLS("./cert.pem", "./key.pem"); err != nil {
		log.Fatal(err)
	}

}

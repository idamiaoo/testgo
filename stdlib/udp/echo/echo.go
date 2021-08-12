package main

import (
	"flag"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

const addr = "localhost:4242"

var num int

func init() {
	flag.IntVar(&num, "n", 1, "")
}

func main() {
	flag.Parse()

	go func() { log.Fatal(echoServer()) }()

	var eg errgroup.Group
	for i := 0; i < num; i++ {
		eg.Go(clientMain)

	}
	_ = eg.Wait()
}

func echoServer() error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			return err
		}

		log.Println("server receive:", string(data[:n]))

		response := remoteAddr.String() + ":" + string(data[:n])

		if _, err := conn.WriteToUDP([]byte(response), remoteAddr); err != nil {
			return err
		}
	}
}

func clientMain() error {
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	for {
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			return err
		}
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			return err
		}
		log.Println("client receive:", string(data[:n]))
	}
}

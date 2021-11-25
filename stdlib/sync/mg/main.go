package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	go func() {
		for {
			<-ch1
			fmt.Println(1)
			<-time.After(1 * time.Second)
			ch2 <- 1
		}

	}()

	go func() {
		for {
			<-ch2
			fmt.Println(2)
			<-time.After(1 * time.Second)
			ch3 <- 1
		}
	}()

	go func() {
		for {
			<-ch3
			fmt.Println(3)
			<-time.After(1 * time.Second)
			ch1 <- 1
		}
	}()

	ch1 <- 1
	select {}
}

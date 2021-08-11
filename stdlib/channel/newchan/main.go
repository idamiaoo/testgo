package main

import "fmt"

func main() {
	ch := new(chan struct{})

	go func() {
		*ch <- struct{}{}
	}()
	<-*ch
	fmt.Println("success")
}

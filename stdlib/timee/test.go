package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	<-time.After(time.Duration(20) * time.Second)
	d := time.Since(now)
	fmt.Println(d)
	fmt.Println(time.Now())
}

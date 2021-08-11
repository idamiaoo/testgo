package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func main() {
	s := "123456789"
	r := bytes.NewReader([]byte(s))
	rr := bufio.NewReader(r)
	fmt.Println(rr.Peek(1))
	rr.Discard(1)
	fmt.Println(rr.Peek(1))
	fmt.Println(rr.Peek(1))

	vv := make([]byte, 4)
	io.ReadFull(rr, vv)
	fmt.Println(string(vv))

}

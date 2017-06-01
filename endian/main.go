package main

import (
	"fmt"
	"unsafe"
)

func main() {
	datas := 0x01020304
	ptr := (*[4]byte)(unsafe.Pointer(&datas))

	for i, v := range *ptr {
		fmt.Println(i, " : ", v, &ptr[i])
	}
}

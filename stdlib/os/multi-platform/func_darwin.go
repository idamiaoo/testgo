package main

import (
	"fmt"
	"runtime"
)

func Print() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

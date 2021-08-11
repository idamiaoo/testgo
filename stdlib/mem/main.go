package main

import (
	"bufio"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func f() {
	container := make([]*bufio.Reader, 8)
	log.Println("> loop.")
	// slice会动态扩容，用它来做堆内存的申请
	for i := 0; i < 512*1024; i++ {
		reader := bufio.NewReaderSize(nil, 10)
		container = append(container, reader)
	}
	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}

func main() {
	log.Println("start.")
	f()

	log.Println("force gc.")
	runtime.GC() // 调用强制gc函数

	log.Println("done.")
	http.ListenAndServe("0.0.0.0:10001", nil)
}

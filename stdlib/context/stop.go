package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())

	for i := 0; i < 10; i++ {
		go func(ctx context.Context, i int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("[%d] 监控退出，停止了...\n", i)
					return
				default:
					fmt.Printf("[%d] goroutine监控中...\n", i)
					time.Sleep(2 * time.Second)
				}
			}
		}(ctx, i)
	}

	time.AfterFunc(10*time.Second, func() {
		fmt.Println("可以了，通知监控停止")
		cancel()
	})

	for {
	}
}

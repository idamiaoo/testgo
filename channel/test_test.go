package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			n, ok := <-ch
			if !ok {
				fmt.Println("n exit")
				return
			}
			fmt.Println("n:", n)
		}

	}()
	for {
		n, ok := <-ch
		if !ok {
			fmt.Println("nn exit")
			break
		}
		fmt.Println("nn:", n)
	}

	<-time.After(5 * time.Second)

}

func TestClose(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		<-time.After(5 * time.Second)
		close(ch)
	}()
	<-ch
}

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

type CLocker struct {
	ch chan struct{}
}

func (locker *CLocker) Lock() {
	<-locker.ch
}

func (locker *CLocker) UnLock() {
	locker.ch <- struct{}{}
}

func TestChanLock(t *testing.T) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	l := &CLocker{
		ch: ch,
	}

	go func() {
		t.Log("g1 try hold lock")
		l.Lock()
		t.Log("g1 hold lock")
		time.Sleep(2 * time.Second)
		l.UnLock()
		t.Log("g1 release lock")
	}()

	go func() {
		t.Log("g2 try hold lock")
		l.Lock()
		t.Log("g2 hold lock")
		time.Sleep(3 * time.Second)
		l.UnLock()
		t.Log("g2 release lock")

	}()

	time.Sleep(10 * time.Second)
}

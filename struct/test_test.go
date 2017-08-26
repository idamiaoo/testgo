package structt

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type t1 struct {
	count int
}

type t2 struct {
	count int
}

type t3 struct {
	a *t1
	b *t2
}

var base = t3{
	a: &t1{},
	b: &t2{},
}

func getbase() t3 {
	return base
}

func TestPS(t *testing.T) {
	go func() {
		b := getbase()
		for {
			<-time.After(2 * time.Second)
			fmt.Println(b.a.count, b.b.count)
		}
	}()
	b := getbase()
	for {
		<-time.After(5 * time.Second)
		b.a.count++
		b.b.count++
	}
}

func TestFmt(t *testing.T) {
	fmt.Printf("%.2f\n", 10.2)
}

func TestJoin(t *testing.T) {
	fmt.Println(strings.Join([]string{"aa", "bb", "cc"}, ","))
}

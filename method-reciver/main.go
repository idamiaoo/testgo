package main

import "fmt"

// Friend .
type Friend interface {
	SayHello(string)
	SetName(string)
}

type T struct {
	name string
}

func (t T) SayHello(name string) {
	fmt.Printf("[t]%s: hello %s\n", t.name, name)
}

func (t T) SetName(name string) {
	t.name = name
}

type P struct {
	name string
}

func (p *P) SayHello(name string) {
	fmt.Printf("[p]%s: hello %s\n", p.name, name)
}

func (p *P) SetName(name string) {
	p.name = name
}

var _ Friend = T{}
var _ Friend = &T{}

// var _ Friend = P{}
var _ Friend = &P{}

type People struct {
	Name string
	Age  int
}

func (p *People) Grow() {
	p.Age++
}

func main() {
	t1 := T{}
	t1.SayHello("carl")
	t2 := &T{}
	t2.SayHello("carl")

	p1 := P{}
	p1.SayHello("carl")
	p2 := P{}
	p2.SayHello("carl")

	t1.SetName("t1")
	t1.SayHello("carl")
	t2.SetName("t2")
	t2.SayHello("carl")

	p1.SetName("p1")
	p1.SayHello("carl")

	p := People{
		Name: "carl",
		Age:  29,
	}
	p.Grow()
	fmt.Println(p)
}

package mmap

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	d := make(map[float64]string)
	d[1.1] = "addd"
	fmt.Println(d[1.1])
}

func TestTTTT(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a[0:1])
}

func TestMapS(t *testing.T) {
	a := make(map[uint32]map[int64][]string)
	a[1] = make(map[int64][]string)
	a1 := a[1]
	a1[1] = append(a1[1], "bohler")
	fmt.Println(a)
}

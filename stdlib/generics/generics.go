package main

import (
	"fmt"
	"sort"
)

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Comparable interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64
}

func Min[T Comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type SortAble interface {
	any
	Less(SortAble) bool
}

func Sort[T any](v []T, less func(a, b T) bool) {
	sort.SliceStable(v, func(i, j int) bool {
		return less(v[i], v[j])
	})
}

func Sort2[T SortAble](v []T) {
	sort.SliceStable(v, func(i, j int) bool {
		return v[i].Less(v[j])
	})
}

type S struct {
	a int
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	fmt.Println()
	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Println(Min(1, 3))
	fmt.Println(Min(0.1, 0.3))
	v1 := []int{1, 3, 2, 9, 0}
	Sort(v1, func(a, b int) bool { return a < b })
	fmt.Println(v1)
	v2 := []S{
		{a: 9},
		{a: -1},
	}
	Sort(v2, func(a, b S) bool { return a.a < b.a })
	fmt.Println(v2)
}

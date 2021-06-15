package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("d1")
	}()

	defer func() {
		fmt.Println("d2")
	}()

	defer func() {
		fmt.Println("d3")
		panic("panic2")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("panic")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("begin")
	defer fmt.Println("over")
	os.Exit(0)
}

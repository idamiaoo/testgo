package bibao

import (
	"log"
	"testing"
)

func getID() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestGetID(t *testing.T) {
	f1 := getID()

	for i := 0; i < 10; i++ {
		log.Println(f1())
	}

	f2 := getID()
	for i := 0; i < 10; i++ {
		log.Println(f2())
	}
}

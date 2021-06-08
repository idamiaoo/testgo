package bibao

import (
	"fmt"
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

type ud struct {
	a int
	b int
}

func TestNil(t *testing.T) {
	var u *ud
	fmt.Println(u)
	if u == nil {
		u = &ud{
			a: 1,
			b: 2,
		}
	}
	fmt.Println(u)
}

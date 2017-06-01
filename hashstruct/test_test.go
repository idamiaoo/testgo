package hashstruct

import (
	"log"
	"testing"

	"github.com/mitchellh/hashstructure"
)

type ss struct {
	Phone string
	email string
	XX    []int
}

type SS struct {
	Name string
	Age  int
	Ex   *ss
}

func TestStruct(t *testing.T) {
	v := &SS{
		Name: "bohler",
		Age:  28,
		Ex: &ss{
			Phone: "1380225041022",
			email: "@163333",
			XX:    []int{1, 2, 3},
		},
	}
	hash, err := hashstructure.Hash(v, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(hash)
}

//17278580130336935856

func testdot(ints ...int) {
	println(len(ints))
}

func TestDot(t *testing.T) {
	testdot(1, 2, 3, 4, 5, 6, 7)
}

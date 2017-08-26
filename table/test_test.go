package table

import (
	"testing"

	"fmt"

	"github.com/modood/table"
)

type User struct {
	Name string
	Age  int
}

func TestTable(t *testing.T) {
	u := User{
		Name: "aaaaaaaaaaaaaaaaaaaaaaaa",
		Age:  1,
	}
	n := table.Table([]User{u})
	fmt.Println(n)
}

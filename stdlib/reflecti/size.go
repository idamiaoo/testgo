package reflecti

import (
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.TypeOf(int32(0))
	fmt.Println(typ.Size())
}

package reflecti

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRef(t *testing.T) {
	b := Bohler{
		Name:  "holly",
		age:   26,
		Phone: "13802250410",
		Addr:  "guiyang",
	}

	Value := reflect.ValueOf(&b)
	Type := reflect.Indirect(Value).Type()
	fmt.Println(Type.Name())
	fmt.Println(Type.Kind())
	fmt.Println(Type.NumField())
	fmt.Println(Value.String())
	fmt.Println(reflect.Indirect(Value).Type().Name())
	fmt.Println(Type.Field(1).Type.Name())
	fmt.Println(Type.NumMethod())
	fmt.Println(Type.Method(0).Name, Type.Method(0).Type)

	fmt.Println(Value.Elem().Field(0).String())
}

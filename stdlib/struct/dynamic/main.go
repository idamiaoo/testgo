package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Age    int
	Gender int
}

type Teacher struct {
	User
	Title   string
	Subject string
}

type Teacher2 struct {
	User    User
	Title   string
	Subject string
}

func main() {
	t1 := &Teacher{
		User: User{
			Name:   "carl",
			Age:    30,
			Gender: 1,
		},
		Title:   "大师",
		Subject: "语文",
	}

	t2 := &Teacher2{
		User: User{
			Name:   "carl",
			Age:    30,
			Gender: 1,
		},
		Title:   "大师",
		Subject: "语文",
	}

	v1 := reflect.ValueOf(t1)
	v1 = reflect.Indirect(v1)
	for i := 0; i < v1.NumField(); i++ {
		fmt.Println(v1.Type().Field(i).Name, v1.Type().Field(i).Anonymous)
	}
	fmt.Println("-------------------------------")
	v2 := reflect.ValueOf(t2)
	v2 = reflect.Indirect(v2)
	for i := 0; i < v2.NumField(); i++ {
		fmt.Println(v2.Type().Field(i).Name, v2.Type().Field(i).Anonymous)
	}
}

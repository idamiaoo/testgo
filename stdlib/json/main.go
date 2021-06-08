package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Type int

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Hello struct {
	Content string    `json:"content"`
	At      time.Time `json:"time,string"`
	Ca      Type      `json:"ca"`
	Info
}

func main() {
	jsonstr := `{"name":"holly","age":10}`
	var h Hello
	if err := json.Unmarshal([]byte(jsonstr), &h); err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)

	zh := Hello{
		Info: Info{
			Name: "holly",
			Age:  10,
		},
		Content: "hello",
		At:      time.Now(),
		Ca:      8,
	}
	str, err := json.Marshal(zh)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))

}

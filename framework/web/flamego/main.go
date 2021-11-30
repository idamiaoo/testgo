package main

import (
	"fmt"
	"strings"

	"github.com/flamego/flamego"
)

func main() {
	f := flamego.New()
	f.Get("/users/{name}", func(c flamego.Context) string {
		fmt.Println(c.Params())
		return fmt.Sprintf("The user is %s", c.Param("name"))
	})
	f.Get("/posts/{year}-{month}-{day}.html", func(c flamego.Context) string {
		fmt.Println(c.Params())
		return fmt.Sprintf(
			"The post date is %d-%d-%d",
			c.ParamInt("year"), c.ParamInt("month"), c.ParamInt("day"),
		)
	})
	f.Get("/geo/{state}/{city}", func(c flamego.Context) string {
		fmt.Println(c.Params())
		return fmt.Sprintf(
			"Welcome to %s, %s!",
			strings.Title(c.Param("city")),
			strings.ToUpper(c.Param("state")),
		)
	})
	f.Run()
}

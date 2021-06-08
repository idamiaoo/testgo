package main

import (
	"github.com/kataras/iris/v12"
)

var P struct {
	Age  int
	Name string
}

func main() {
	P.Age = 26
	P.Name = "bohler"
	app := iris.New()
	tmpl := iris.HTML("./views", ".html")
	tmpl.Delims("{{", "}}")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.ViewData("P", P)
		ctx.View("index.html")
	})
	app.Listen(":8180")
}

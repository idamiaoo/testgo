package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/view"
)

var P struct {
	Age  int
	Name string
}

func main() {
	P.Age = 26
	P.Name = "bohler"
	app := iris.New()
	app.AttachView(view.HTML("./views", ".html").Reload(true))
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.ViewData("P", P)
		ctx.View("index.html")
	})
	app.Run(iris.Addr(":8180"), iris.WithCharset("UTF-8"))
}

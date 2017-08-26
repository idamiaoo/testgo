package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

type P struct {
	Age  int
	Name string
}

func main() {
	var p P
	p.Age = 20
	p.Name = "bohler"
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"P": &p,
		})
	})
	router.Run(":8080")
}

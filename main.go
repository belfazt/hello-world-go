package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World!"})
	})

	app.Handle("GET", "/status", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": "ok"})
	})

	app.Run(iris.Addr(":8080"))
}

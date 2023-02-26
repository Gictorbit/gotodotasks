package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("webserver/templates", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/login", iris.StatusFound)
	})
	app.Get("/login", showLogin)
	app.Get("/taskmanager", showTaskManager)
	app.HandleDir("/static", "./webserver/templates/static")
	app.Listen(":2030")
}

func showLogin(ctx iris.Context) {
	if err := ctx.View("login.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

func showTaskManager(ctx iris.Context) {
	if err := ctx.View("taskmanager.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

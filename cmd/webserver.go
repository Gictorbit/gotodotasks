package main

import (
	"github.com/kataras/iris/v12"
)

func NewWebServer() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel("debug")
	// Register the HTML view engine
	app.RegisterView(iris.HTML("templates", ".html"))

	// Define the route handlers
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/login", iris.StatusFound)
	})

	app.Get("/login", showLogin)
	app.Get("/taskmanager", showTaskManager)

	// Serve the static files
	app.HandleDir("/assets", "templates/assets")
	app.HandleDir("/css", "templates/css")
	app.HandleDir("/js", "templates/js")

	return app
}

func showLogin(ctx iris.Context) {
	ctx.ViewData("domain", Domain)
	if err := ctx.View("login.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

func showTaskManager(ctx iris.Context) {
	ctx.ViewData("domain", Domain)
	if err := ctx.View("taskmanager.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

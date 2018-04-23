package main

import (
	"github.com/go-macaron/binding"
	"github.com/go-macaron/jade"

	"github.com/PhillBeck/ingresse-backend/conf"
	"github.com/PhillBeck/ingresse-backend/handler"
	"github.com/PhillBeck/ingresse-backend/model"
	macaron "gopkg.in/macaron.v1"
)

func main() {
	app := macaron.New()
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
	}))
	app.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "public/templates",
	}))

	setupRoutes(app)

	app.Run(conf.GetHttpPort())
}

func setupRoutes(app *macaron.Macaron) {
	userHandler := handler.NewUserHandler()

	app.Group("/users", func() {
		app.Post("/", binding.Bind(model.User{}), userHandler.Post)
		app.Get("/:ID", userHandler.GetOne)
		app.Put("/:ID", binding.Bind(model.User{}), userHandler.Put)
		app.Delete("/:ID", userHandler.Delete)
	})
}

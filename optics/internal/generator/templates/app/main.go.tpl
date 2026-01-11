package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	{{range .Modules}}
	{{.}}.Register(app)
	{{end}}

	app.Listen(":8080")
}

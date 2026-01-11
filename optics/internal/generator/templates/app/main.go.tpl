package main

import (
	"github.com/gofiber/fiber/v2"
	{{- if .Modules }}

	{{- range .Modules }}
	"{{ . }}"
	{{- end }}

	{{- end }}
)

func main() {
	app := fiber.New()

	{{- if .Modules }}
	{{- range .Modules }}
	{{ . }}.Register(app)
	{{- end }}
	{{- end }}

	app.Listen(":8080")
}

package {{.Name}}

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, controller *{{.Pascal}}Controller) {
	group := app.Group("/{{.Name}}")

	group.Get("/", controller.GetAll)
	group.Get("/:id", controller.GetByID)
	group.Post("/", controller.Create)
	group.Put("/:id", controller.Update)
	group.Delete("/:id", controller.Delete)
}

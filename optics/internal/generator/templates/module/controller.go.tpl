package {{.Name}}

import "github.com/gofiber/fiber/v2"

type {{.Pascal}}Controller struct {
	service *{{.Pascal}}Service
}

func New{{.Pascal}}Controller(service *{{.Pascal}}Service) *{{.Pascal}}Controller {
	return &{{.Pascal}}Controller{service: service}
}

func (c *{{.Pascal}}Controller) GetAll(ctx *fiber.Ctx) error {
	return ctx.SendString("Get all {{.Name}}")
}

func (c *{{.Pascal}}Controller) GetByID(ctx *fiber.Ctx) error {
	return ctx.SendString("Get {{.Name}} by id")
}

func (c *{{.Pascal}}Controller) Create(ctx *fiber.Ctx) error {
	return ctx.SendString("Create {{.Name}}")
}

func (c *{{.Pascal}}Controller) Update(ctx *fiber.Ctx) error {
	return ctx.SendString("Update {{.Name}}")
}

func (c *{{.Pascal}}Controller) Delete(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete {{.Name}}")
}

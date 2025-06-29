package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/controllers/svc1"
	"go-fiber/controllers/svc2"
)

func Routes() *fiber.App {
	app := fiber.New()

	app_svc1 := app.Group("/svc1")
	app_svc1.Get("/req1", svc1.Req1)
	app_svc1.Get("/req2", svc1.Req2)

	app_svc2 := app.Group("/svc2")
	app_svc2.Get("/req1", svc2.Req1)
	app_svc2.Get("/req2", svc2.Req2)

	return app
}

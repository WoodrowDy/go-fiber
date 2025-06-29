package svc1

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/models"
)

func Req1(c *fiber.Ctx) error {
	result := models.GetAdminList()
	return c.JSON(result)
}

func Req2(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"svc1": "req2"})
}

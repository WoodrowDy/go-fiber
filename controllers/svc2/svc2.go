package svc2

import "github.com/gofiber/fiber/v2"

func Req1(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"svc2": "req1"})
}

func Req2(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"svc2": "req2"})
}

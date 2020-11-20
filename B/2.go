package B

import "github.com/gofiber/fiber/v2"

func two(ctx *fiber.Ctx) error {
	return ctx.SendString("B, 2 👋!")
}

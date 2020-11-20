package B

import "github.com/gofiber/fiber/v2"

func three(ctx *fiber.Ctx) error {
	return ctx.SendString("B, 3 👋!")
}

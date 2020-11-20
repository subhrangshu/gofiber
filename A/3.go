package A

import "github.com/gofiber/fiber/v2"

func three(ctx *fiber.Ctx) error {
	return ctx.SendString("A, 3 👋!")
}

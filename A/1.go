package A

import "github.com/gofiber/fiber/v2"

func one(ctx *fiber.Ctx) error {
	return ctx.SendString("A, 1 👋!")
}

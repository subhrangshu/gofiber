package C

import "github.com/gofiber/fiber/v2"

func three(ctx *fiber.Ctx) error {
	return ctx.SendString("C, 3 👋!")
}

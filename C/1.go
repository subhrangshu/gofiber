package C

import "github.com/gofiber/fiber/v2"

func one(ctx *fiber.Ctx) error {
	return ctx.SendString("C, 1 👋!")
}

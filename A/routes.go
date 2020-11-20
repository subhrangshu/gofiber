package A

import "github.com/gofiber/fiber/v2"

func Routes(router fiber.Router) {
	router.Get("/1", one)
	router.Get("/2", two)
	router.Get("/3", three)
}

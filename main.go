package main

import (
	"fibr/A"
	"fibr/B"
	"fibr/C"
	"fibr/root"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/qinains/fastergoding"
	"log"
)

func main() {
	fastergoding.Run()
	// Fiber instance
	app := fiber.New()
	app.Use(compress.New())

	// Routes
	//app.Get("/", hello)
	SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

func SetupRoutes(router fiber.Router) {
	router.Get("/", root.Hello)
	A.Routes(router.Group("A"))
	B.Routes(router.Group("B"))
	C.Routes(router.Group("C"))
}

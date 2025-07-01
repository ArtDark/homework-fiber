package main

import (
	"publicator/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	pub := fiber.New()

	pub.Use(recover.New())

	pages.NewHandler(pub)

	pub.Listen(":9001")

}

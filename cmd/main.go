package main

import (
	"publicator/config"
	"publicator/internal/pages"
	"publicator/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
		"github.com/gofiber/fiber/v2/middleware/recover"
	slfiber "github.com/samber/slog-fiber"
)

func main() {
	
	config.Init()

	cfg := config.NewMainConfig()

	engine := html.New("./templates", ".html") 

	pub := fiber.New(fiber.Config{
		Views: engine,
	})

	pub.Use(recover.New())
	pub.Use(slfiber.New(logger.NewService(cfg.Log).Info))
	pub.Use(slfiber.New(logger.NewService(cfg.Log).Error))

	pages.NewHandler(pub)

	pub.Listen(":9001")

}

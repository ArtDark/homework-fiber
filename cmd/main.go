package main

import (
	"homework-fiber/config"
	"homework-fiber/internal/pages"
	"homework-fiber/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slfiber "github.com/samber/slog-fiber"
)

func main() {

	config.Init()

	cfg := config.NewMainConfig()

	pub := fiber.New()

	pub.Use(recover.New())
	pub.Use(slfiber.New(logger.NewService(cfg.Log).Info))
	pub.Use(slfiber.New(logger.NewService(cfg.Log).Error))

	pub.Static("/public", "./public")

	pages.NewHandler(pub)

	
	pub.Listen(":9001")

}

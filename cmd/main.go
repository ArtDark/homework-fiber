package main

import (
	"homework-fiber/config"
	"homework-fiber/internal/pages"
	"homework-fiber/pkg/database"
	"homework-fiber/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slfiber "github.com/samber/slog-fiber"
)

func main() {

	config.Init()

	cfg := config.NewMainConfig()

	pub := fiber.New()

	logService := logger.NewService(cfg.Log)

	pub.Use(recover.New())
	pub.Use(slfiber.New(logService.Info))
	pub.Use(slfiber.New(logService.Error))


	dbPool := database.CreateDbPool(cfg.Database,logService)
	defer database.CloseDbPool(dbPool,logService)

	pub.Static("/public", "./public")

	pages.NewHandler(pub)

	pub.Listen(":9001")

}

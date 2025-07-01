package pages

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{router: router}
	h.router.Group("api/v1")
	h.router.Get("/", h.home)

}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	return c.SendString("Hello, World")
}

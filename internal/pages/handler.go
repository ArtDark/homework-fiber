package pages

import "github.com/gofiber/fiber/v2"

type Handler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &Handler{router: router}
	apiV1 := h.router.Group("api/v1")
	apiV1.Get("/", h.home)

}

func (h *Handler) home(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	return c.SendString("Hello, World")
}

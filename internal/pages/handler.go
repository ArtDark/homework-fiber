package pages

import (
	"homework-fiber/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type Handler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &Handler{router: router}

	h.router.Get("/", h.home)

	api := h.router.Group("/api")
	api.Get("/register", h.register)
}

func (h *Handler) home(c *fiber.Ctx) error {
	main := views.Main()

	return httpAdaptor(c, main)

}

func (h *Handler) register(c *fiber.Ctx) error {
	return httpAdaptor(c, views.Register())
}

func httpAdaptor(c *fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}

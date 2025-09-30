package pages

import (
	"fmt"
	"homework-fiber/views"
	"homework-fiber/views/components"

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
	api.Get("/register", h.registerPage)

	api.Post("/register", h.register)
}

func (h *Handler) home(c *fiber.Ctx) error {
	main := views.Main()

	return httpAdaptor(c, main)

}

func (h *Handler) registerPage(c *fiber.Ctx) error {
	return httpAdaptor(c, views.RegisterPage())
}

func (h *Handler) register(c *fiber.Ctx) error {

	form := struct {
		Login    string `form:"login"`
		Password string `form:"password"`
		Email    string `form:"email"`
	}{}

	c.BodyParser(&form)

	fmt.Println(form)

	return httpAdaptor(c, components.Notifucation(components.NotificationProps{
		TypeMsg: components.NotificationTypeSuccess,
		Message: "Success",
	}))
}

func httpAdaptor(c *fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}

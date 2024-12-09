package routes

import (
	"exclusive-sample/internal/api/controller"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(r fiber.Router, adminController *controller.Admin) {
	r.Post("/", adminController.CreateAdmin)
	r.Get("/", adminController.GetAllAdmin)
}

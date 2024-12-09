package routes

import (
	"exclusive-sample/internal/api/controller"

	"github.com/gofiber/fiber/v2"
)

func VendorRoutes(r fiber.Router, vendorController *controller.Vendor) {
	r.Post("/", vendorController.CreateVendor)
	r.Get("/", vendorController.GetVendors)
}

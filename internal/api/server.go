package http

import (
	"exclusive-sample/internal/api/controller"
	"exclusive-sample/internal/api/routes"
	"exclusive-sample/internal/config"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ServerHTTP struct {
	app *fiber.App
}

func NewServerHTTP(vendorHandler *controller.Vendor, adminHandler *controller.Admin) *ServerHTTP {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	vendorGroup := app.Group("/vendor")
	routes.VendorRoutes(vendorGroup, vendorHandler)

	adminGroup := app.Group("/admin")
	routes.AdminRoutes(adminGroup, adminHandler)

	return &ServerHTTP{app: app}
}

func (sh *ServerHTTP) Start(cfg config.Config, infoLog *log.Logger, errorLog *log.Logger) {
	err := sh.app.Listen(fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		errorLog.Fatal(err)
	}
}

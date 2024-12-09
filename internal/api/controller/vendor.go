package controller

import (
	"exclusive-sample/internal/models"
	"exclusive-sample/internal/service"
	"exclusive-sample/internal/utils/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Vendor struct {
	service *service.Vendor
}

func NewVendorController(service *service.Vendor) *Vendor {
	return &Vendor{service: service}
}

func (c *Vendor) CreateVendor(ctx *fiber.Ctx) error {
	var vendor models.Vendor

	if err := ctx.BodyParser(&vendor); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Invalid input data: " + err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	err := validator.New().Struct(vendor)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Constraints not statisfied: " + err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	createvendor, err := c.service.CreateVendor(&vendor)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Failed to create vendor: " + err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(responses.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data: &fiber.Map{
			"vendor": createvendor,
		},
	})

}

func (c *Vendor) GetVendors(ctx *fiber.Ctx) error {
	vendor, err := c.service.GetAllVendors()
	if err != nil || vendor == nil {
		return ctx.Status(http.StatusNotFound).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: "vendors not found" + err.Error(),
			ErrorTime:    time.Now(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(responses.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data: &fiber.Map{
			"Vendors": vendor,
		},
	})
}

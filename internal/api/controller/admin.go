package controller

import (
	"exclusive-sample/internal/models"
	"exclusive-sample/internal/service"
	"exclusive-sample/utils/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Admin struct {
	service *service.Admin
}

func NewAdminController(service *service.Admin) *Admin {
	return &Admin{service: service}
}

func (c *Admin) CreateAdmin(ctx *fiber.Ctx) error {
	var admin models.Admin

	if err := ctx.BodyParser(&admin); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Invalid input data: " + err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	err := validator.New().Struct(admin)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Constraints not statisfied: " + err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	createdAdmin, err := c.service.CreateAdmin(&admin)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Failed to create admin: " + err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(responses.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data: &fiber.Map{
			"admin": createdAdmin,
		},
	})
}

func (c *Admin) GetAllAdmin(ctx *fiber.Ctx) error {
	admin, err := c.service.GetAllAdmin()
	if err != nil || admin == nil {
		return ctx.Status(http.StatusNotFound).JSON(responses.ErrorResponse{
			ApiPath:      ctx.OriginalURL(),
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: "admin not found" + err.Error(),
			ErrorTime:    time.Now(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(responses.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data: &fiber.Map{
			"admin": admin,
		},
	})
}

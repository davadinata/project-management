package controller

import (
	"github.com/davadinata/project-management/models"
	"github.com/davadinata/project-management/services"
	"github.com/davadinata/project-management/utils"
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.Bind().Body(user); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Failed to register user", err.Error())
	}
	return utils.Success(ctx, "User registered successfully", user)
}

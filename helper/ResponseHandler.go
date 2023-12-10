package helper

import "github.com/gofiber/fiber/v2"

func ResponseHandler(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func UnAuthorizedResponse(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": message,
	})
}

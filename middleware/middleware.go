package middleware

import (
	"github.com/go-fiber-crud/helper"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Auth(ctx *fiber.Ctx) error {
	AuthHeader := ctx.Get("Authorization")

	if AuthHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	authParts := strings.SplitN(AuthHeader, " ", 2)
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token := authParts[1]

	_, err := helper.VerifyJwtToken(token)
	claims, err := helper.DecodeJwtToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	role := claims["role"].(string)

	if role != "admin" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Forbidden access",
		})
	}

	//ctx.Locals("userInfo", claims)
	//ctx.Locals("role", claims["role"])

	return ctx.Next()
}

package handle

import "github.com/gofiber/fiber/v2"

func SendToClientError(ctx *fiber.Ctx, errMsg string, statusCode int) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"message": errMsg,
	})
}

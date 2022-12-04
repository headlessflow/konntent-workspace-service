package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthorizeMobilisim() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

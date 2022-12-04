package middleware

import (
	"konntent-workspace-service/pkg/claimer"
	"konntent-workspace-service/pkg/middlewarepkg"
	"konntent-workspace-service/pkg/utils"
	"konntent-workspace-service/pkg/utils/userutil"

	"github.com/gofiber/fiber/v2"
)

func Authorize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt := c.Context().Value(utils.Claimer).(claimer.Claimer)
		jwtJson, valid := jwt.IsValid(c.UserContext(), middlewarepkg.GetAuthorizationHeader(c))

		if !valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		userutil.PutAuthModel(c, jwtJson)

		return c.Next()
	}
}

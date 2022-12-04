package userutil

import (
	"context"
	"encoding/json"
	"konntent-workspace-service/pkg/claimer"
	"konntent-workspace-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUserModel(c context.Context) *claimer.Model {
	jwt := c.Value(utils.Claimer).(claimer.Claimer)
	model := jwt.GetModel(c)

	return model
}

func PutAuthModel(c *fiber.Ctx, input []byte) {
	var model claimer.Model
	_ = json.Unmarshal(input, &model)

	c.Locals(utils.AuthCtx, &model)
}

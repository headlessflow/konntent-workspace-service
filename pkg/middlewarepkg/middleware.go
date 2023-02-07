package middlewarepkg

import (
	"konntent-workspace-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func PutHeaders(ctx *fiber.Ctx) error {
	var headerMap = utils.ContextHeader{}

	_ = ctx.ReqHeaderParser(&headerMap)
	ctx.Locals(utils.HeaderMapCtx, &headerMap)

	return ctx.Next()
}

func GetAuthorizationHeader(ctx *fiber.Ctx) []byte {
	var headerMap = ctx.Locals(utils.HeaderMapCtx).(map[string]string)

	if val, ok := headerMap[utils.HeaderAuthorization]; ok {
		return []byte(val)
	}

	return []byte{}
}

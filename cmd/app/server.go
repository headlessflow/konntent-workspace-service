package main

import (
	di "konntent-workspace-service"
	"konntent-workspace-service/internal/app"
	"konntent-workspace-service/internal/app/middleware"
	"konntent-workspace-service/pkg/claimer"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/middlewarepkg"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/pg"
	"konntent-workspace-service/pkg/utils"
	"konntent-workspace-service/pkg/validation"

	recoverpkg "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	logger     *zap.Logger
	pgInstance pg.Instance

	jwtInstance claimer.Claimer
	nrInstance  nrclient.NewRelicInstance
}

func initServer(sv *server) *fiber.App {
	fApp := fiber.New(fiber.Config{
		BodyLimit: constants.AppRequestBodyLimit,
	})
	fApp.Use(recoverpkg.New(recoverpkg.Config{
		EnableStackTrace: true,
	}))

	sv.initCommonMiddlewares(fApp)

	route := di.InitAll(
		sv.logger,
		sv.pgInstance,
		sv.nrInstance,
	)
	route.SetupRoutes(&app.RouteCtx{
		App: fApp,
	})

	return fApp
}

func initLogger() *zap.Logger {
	zc := zap.NewDevelopmentEncoderConfig()
	zc.EncodeLevel = zapcore.CapitalColorLevelEncoder

	l := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zc),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))
	return l
}

func (s *server) initCommonMiddlewares(app *fiber.App) {
	validator := validation.InitValidator()
	app.Use(middleware.NewRelicMiddleware(s.nrInstance))
	app.Use(func(c *fiber.Ctx) error {
		c.Locals(utils.Validator, validator)
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Locals(utils.Claimer, s.jwtInstance)
		return c.Next()
	})

	app.Use(middlewarepkg.PutHeaders)
}

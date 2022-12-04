package main

import (
	di "konntent-workspace-service"
	"konntent-workspace-service/internal/app"
	"konntent-workspace-service/internal/app/middleware"
	"konntent-workspace-service/pkg/claimer"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/dummyclient"
	"konntent-workspace-service/pkg/middlewarepkg"
	"konntent-workspace-service/pkg/nrclient"
	"konntent-workspace-service/pkg/rabbit"
	"konntent-workspace-service/pkg/utils"
	"konntent-workspace-service/pkg/validation"

	recoverpkg "github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type server struct {
	logger      *logrus.Logger
	dummyClient dummyclient.Client
	mqProducer  rabbit.Client
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
		sv.dummyClient,
		sv.mqProducer,
		sv.jwtInstance,
		sv.nrInstance,
	)
	route.SetupRoutes(&app.RouteCtx{
		App: fApp,
	})

	return fApp
}

func initLogger() *logrus.Logger {
	return logrus.New()
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

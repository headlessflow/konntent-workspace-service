package app

import (
	"konntent-workspace-service/internal/app/handler"
	"konntent-workspace-service/internal/app/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteCtx struct {
	App *fiber.App
}

type Router interface {
	SetupRoutes(r *RouteCtx)
}

type route struct {
	dummyHandler handler.DummyHandler
}

func NewRoute(dummyHandler handler.DummyHandler) Router {
	return &route{
		dummyHandler: dummyHandler,
	}
}

func (r *route) SetupRoutes(rc *RouteCtx) {
	rc.App.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	v1 := rc.App.Group("/v1")

	r.dummyRoutes(v1)
}

func (r *route) dummyRoutes(gr fiber.Router) {
	dummy := gr.Group("/dummy", middleware.Authorize(), middleware.AuthorizeMobilisim())

	dummy.Get("/", r.dummyHandler.Endpoint)
}

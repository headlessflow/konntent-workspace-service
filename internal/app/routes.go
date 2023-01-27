package app

import (
	"konntent-workspace-service/internal/app/handler"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RouteCtx struct {
	App *fiber.App
}

type Router interface {
	SetupRoutes(r *RouteCtx)
}

type route struct {
	workspaceHandler handler.WorkspaceHandler
}

func NewRoute(workspaceHandler handler.WorkspaceHandler) Router {
	return &route{
        workspaceHandler: workspaceHandler,
	}
}

type Workspace struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (r *route) SetupRoutes(rc *RouteCtx) {
	rc.App.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	v1 := rc.App.Group("/v1")
	v1.Get("/workspaces/:uid", func(ctx *fiber.Ctx) error {
		log.Println("user id: ", ctx.Params("uid"))

		var response = Workspace{
			Id:   1,
			Name: "Netflix",
			Url:  "http://localhost:1234/w/netflix",
		}

		return ctx.JSON(response)
	})

	v1.Post("/workspace", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]interface{}{
			"workspaceId": 1,
		})
	})

	v1.Get("/workspaces", func(ctx *fiber.Ctx) error {
		var response = []Workspace{
			{
				Id:   1,
				Name: "Netflix",
				Url:  "http://localhost:1234/w/netflix",
			},
			{
				Id:   2,
				Name: "Facebook",
				Url:  "http://localhost:1234/w/facebook",
			},
			{
				Id:   3,
				Name: "AirBnb",
				Url:  "http://localhost:1234/w/airbnb",
			},
		}
		return ctx.JSON(response)
	})
}

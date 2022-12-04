package handler

import (
	"konntent-workspace-service/internal/app/orchestration"

	"github.com/gofiber/fiber/v2"
)

type DummyHandler interface {
	Endpoint(c *fiber.Ctx) error
}

type dummyHandler struct {
	dummyOrchestrator orchestration.DummyOrchestrator
}

func NewDummyHandler(do orchestration.DummyOrchestrator) DummyHandler {
	return &dummyHandler{
		dummyOrchestrator: do,
	}
}

func (h *dummyHandler) Endpoint(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

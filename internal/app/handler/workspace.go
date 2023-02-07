package handler

import (
	"konntent-workspace-service/internal/app/dto/request"
	"konntent-workspace-service/internal/app/orchestration"
	"konntent-workspace-service/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

type WorkspaceHandler interface {
	GetWorkspaces(c *fiber.Ctx) error
	GetWorkspace(c *fiber.Ctx) error
	AddWorkspace(c *fiber.Ctx) error
}

type workspaceHandler struct {
	workspaceOrchestration orchestration.WorkspaceOrchestrator
}

func NewWorkspaceHandler(wo orchestration.WorkspaceOrchestrator) WorkspaceHandler {
	return &workspaceHandler{workspaceOrchestration: wo}
}

func (w *workspaceHandler) GetWorkspaces(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		uid = utils.GetUserIDByContext(ctx)
	)

	workspaces, err := w.workspaceOrchestration.GetWorkspaces(ctx, uid)
	if err != nil {
		return c.JSON(nil)
	}

	return c.JSON(workspaces)
}

func (w *workspaceHandler) GetWorkspace(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		req = request.GetWorkspaceRequest{
			UserID: utils.GetUserIDByContext(ctx),
		}
	)

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(nil)
	}

	workspace, err := w.workspaceOrchestration.GetWorkspace(ctx, req)
	if err != nil {
		return c.JSON(nil)
	}

	return c.JSON(workspace)
}

func (w *workspaceHandler) AddWorkspace(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		req = request.AddWorkspaceRequest{
			UserID: utils.GetUserIDByContext(ctx),
		}
	)

	log.Println(req.UserID)

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(nil)
	}

	err := w.workspaceOrchestration.AddWorkspace(ctx, req)
	if err != nil {
		return c.JSON(nil)
	}

	return c.SendStatus(fiber.StatusOK)
}

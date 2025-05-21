package process

import (
	"github.com/gar-id/queued/internal/server/api/types"
	"github.com/gar-id/queued/internal/server/config"
	"github.com/gofiber/fiber/v2"
)

func ConfigDelete(c *fiber.Ctx) (result error) {
	// Parse POST request
	var actionPayload types.ActionPayload
	var statusMessage = make(map[string][]string)
	err := c.BodyParser(&actionPayload)
	if err != nil {
		return c.JSON(General(fiber.StatusBadRequest, "error", c.IP(), err.Error()))
	} else if len(actionPayload.ProgramName) == 0 {
		return c.JSON(General(fiber.StatusBadRequest, "error", c.IP(), "please insert programName input form"))
	}

	// Check process name and execute if exist
	statusMessage, httpStatus, _ := config.DeleteConfig(actionPayload.ProgramName)

	// Return
	if httpStatus == "error" {
		return c.JSON(GeneralObject(fiber.StatusBadRequest, "error", c.IP(), statusMessage))
	}
	return c.JSON(GeneralObject(fiber.StatusOK, "success", c.IP(), statusMessage))
}

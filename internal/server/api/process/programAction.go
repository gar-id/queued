package process

import (
	"fmt"

	"github.com/gar-id/queued/internal/server/api/types"
	"github.com/gar-id/queued/internal/server/config/caches"
	"github.com/gar-id/queued/internal/server/programs"

	"github.com/gofiber/fiber/v2"
)

func ProgramAction(c *fiber.Ctx, action string) (result error) {
	// Parse POST request
	var actionPayload types.ActionPayload
	var statusMessage = make(map[string][]string)
	var httpStatus string
	err := c.BodyParser(&actionPayload)
	if err != nil {
		return c.JSON(General(fiber.StatusBadRequest, "error", c.IP(), err.Error()))
	}

	// Check payload data
	if len(actionPayload.GroupName) > 0 {
		// If groupName payload is more than 0, then do action
		for _, groupName := range actionPayload.GroupName {
			_, ok := caches.Data.GroupIndex[groupName]
			if ok {
				statusMessage[groupName] = append(statusMessage[groupName], programs.Validate(groupName, "group", action)...)
			} else {
				httpStatus = "error"
				statusMessage[groupName] = append(statusMessage[groupName], fmt.Sprintf("group %v not found", groupName))
			}
		}
	}

	if len(actionPayload.ProgramName) > 0 {
		// If programName payload is more than 0, then do action
		for _, programName := range actionPayload.ProgramName {
			_, ok := caches.Data.ProgramConfig[programName]
			if ok {
				statusMessage[programName] = append(statusMessage[programName], programs.Validate(programName, "program", action)...)
			} else {
				httpStatus = "error"
				statusMessage[programName] = append(statusMessage[programName], fmt.Sprintf("program %v not found", programName))
			}
		}
	}

	if len(actionPayload.ProcessName) > 0 {
		// If processName payload is more than 0, then do action
		for _, processName := range actionPayload.ProcessName {
			_, ok := caches.ProcessChannel.Data[processName]
			if ok {
				statusMessage[processName] = append(statusMessage[processName], programs.Validate(processName, "process", action)...)
			} else {
				httpStatus = "error"
				statusMessage[processName] = append(statusMessage[processName], fmt.Sprintf("process %v not found", processName))
			}
		}
	}

	if len(actionPayload.GroupName) == 0 && len(actionPayload.ProgramName) == 0 && len(actionPayload.ProcessName) == 0 {
		return c.JSON(General(fiber.StatusBadRequest, "error", c.IP(), "please insert programName, groupName or processName input form"))
	}

	// Return
	if httpStatus == "error" {
		return c.JSON(GeneralObject(fiber.StatusBadRequest, "error", c.IP(), statusMessage))
	}
	return c.JSON(GeneralObject(fiber.StatusOK, "success", c.IP(), statusMessage))
}

package process

import (
	"fmt"

	"github.com/gar-id/queued/internal/server/config/caches"
	programConfig "github.com/gar-id/queued/internal/server/config/types"

	"github.com/gofiber/fiber/v2"
)

func ProgramLists(c *fiber.Ctx) (result error) {
	// Parse GET queries
	params := c.Queries()

	// Check if programName params is not null and exist
	if params["programName"] != "" {
		_, ok := caches.Data.ProgramConfig[params["programName"]]
		if ok {
			program := make(map[string]programConfig.ProgramConfig)
			program[params["programName"]] = caches.Data.ProgramConfig[params["programName"]]
			return c.JSON(Program(fiber.StatusOK, "success", c.IP(), program))
		} else {
			return c.JSON(General(fiber.StatusBadRequest, "error", c.IP(), fmt.Sprintf("%v not found", params["programName"])))
		}
	}

	// Return all config and status if programName params is null
	return c.JSON(Program(fiber.StatusOK, "success", c.IP(), caches.Data.ProgramConfig))
}

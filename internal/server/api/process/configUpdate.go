package process

import (
	generalCaches "github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/internal/server/config"

	"github.com/gofiber/fiber/v2"
)

func ConfigUpdate(c *fiber.Ctx) (result error) {
	// Reread config in directory
	returnConfig, err := config.RereadConfig(generalCaches.MainConfig.QueueD.Include)
	if err != nil {
		return c.JSON(General(fiber.StatusBadRequest, "error", c.IP(), err.Error()))
	}
	if len(returnConfig) == 0 {
		var statusMessage = make(map[string][]string)
		statusMessage["_queued_status"] = append(statusMessage["_queued_status"], "no update")
		returnConfig = statusMessage
	}

	// Return all config and status if err is null
	return c.JSON(GeneralObject(fiber.StatusOK, "success", c.IP(), returnConfig))
}

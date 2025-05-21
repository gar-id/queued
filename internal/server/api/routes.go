package api

import (
	"github.com/gar-id/queued/internal/server/api/process"
	"github.com/gofiber/fiber/v2"
)

func routes(app *fiber.App) {
	api := app.Group("api/v1")
	apiQueued := api.Group("/queued")
	apiProgram := apiQueued.Group("/program")
	apiConfig := apiQueued.Group("/config")

	// Default response
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(process.Welcome(200, "success", c.IP(), "Welcome to API QueueD"))
	})
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(process.Welcome(200, "success", c.IP(), "Welcome to API QueueD"))
	})
	apiQueued.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(process.Welcome(200, "success", c.IP(), "Welcome to API QueueD"))
	})

	// Program routes
	apiProgram.Get("/", func(c *fiber.Ctx) error {
		return process.ProgramLists(c)
	})
	apiProgram.Post("/stop", func(c *fiber.Ctx) error {
		return process.ProgramAction(c, "stop")
	})
	apiProgram.Post("/start", func(c *fiber.Ctx) error {
		return process.ProgramAction(c, "start")
	})
	apiProgram.Post("/restart", func(c *fiber.Ctx) error {
		return process.ProgramAction(c, "restart")
	})
	apiProgram.Get("/logs", func(c *fiber.Ctx) error {
		return process.TailLogs(c)
	})

	// Config routes
	apiConfig.Get("/", func(c *fiber.Ctx) error {
		return process.ProgramLists(c)
	})
	apiConfig.Post("/add", func(c *fiber.Ctx) error {
		return process.ConfigAdd(c)
	})
	apiConfig.Post("/update", func(c *fiber.Ctx) error {
		return process.ConfigUpdate(c)
	})
	apiConfig.Post("/delete", func(c *fiber.Ctx) error {
		return process.ConfigDelete(c)
	})
}

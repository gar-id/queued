package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/gar-id/queued/internal/general/config/caches"
	"github.com/gar-id/queued/internal/server/api/types"
	"github.com/gar-id/queued/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	// Default error handler
	var customErrorHandler = func(c *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		// Set Content-Type: text/plain; charset=utf-8
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		// Return status code with error message
		var result = types.General{
			HTTP_Code: code,
			Status:    "error",
			ClientIP:  c.IP(),
			Data: struct {
				Date    string "json:\"date\""
				Message string "json:\"message\""
			}{
				Date:    fmt.Sprint(time.Now().String()),
				Message: e.Message,
			}}
		return c.Status(code).JSON(result)
	}

	fibercfg := fiber.Config{
		Prefork:               false,
		ServerHeader:          "QueueD by Gar",
		Concurrency:           256 * 1024 * 30,
		DisableStartupMessage: true,
		ErrorHandler:          customErrorHandler,
	}
	app := fiber.New(fibercfg)

	// Setup cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     tools.DefaultString(caches.MainConfig.QueueD.API.Cors, "*"),
		AllowHeaders:     "Origin, Content-Type, Accept, Cache-Control",
		AllowCredentials: false,
	}))

	// Routes config
	routes(app)

	// Start to listen
	err := app.Listen(caches.MainConfig.QueueD.API.HTTPListen)
	if err != nil {
		tools.ZapLogger("both").Fatal(fmt.Sprintf("Failed to start QueueD RestAPI Server. %v", err))
	}
}

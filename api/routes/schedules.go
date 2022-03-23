package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kl-engineering/schedule/api/handlers"
	"github.com/kl-engineering/schedule/pkg/backend/cmsclient"
)

func SchedulesRouter(app fiber.Router, client cmsclient.Client) {
	app.Get("/schedules/:id", handlers.GetSchedule(client))
}

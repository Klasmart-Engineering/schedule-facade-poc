package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kl-engineering/schedule/api/handlers"
)

func PingRouter(app fiber.Router) {
	app.Get("/ping", handlers.Ping())
}

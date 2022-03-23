package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kl-engineering/schedule/pkg/backend/cmsclient"
)

func GetSchedule(client cmsclient.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		scheduleID := ctx.Params("id")
		result, _ := client.GetSchedule(scheduleID)
		return ctx.JSON(result)
	}
}

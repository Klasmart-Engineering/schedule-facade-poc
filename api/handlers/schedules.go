package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kl-engineering/schedule/pkg/backend/cmsclient"
	"github.com/kl-engineering/schedule/pkg/event/eventpublisher"
)

func GetSchedule(client cmsclient.Client, ev eventpublisher.Publisher) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		scheduleID := ctx.Params("id")
		result, _ := client.GetSchedule(scheduleID)

		// TODO: async publish domain event(s) to messaging (e.g. Kafka).
		//  This is a crude example of 'publishing' the response - whilst not
		//  useful we could emit key events using the facade
		ev.Publish(result)

		return ctx.JSON(result)
	}
}

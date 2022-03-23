package requestlogger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func New(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		logger.Info(
			"request",
			zap.Any("requestid", c.Locals("requestid")),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		return c.Next()
	}
}

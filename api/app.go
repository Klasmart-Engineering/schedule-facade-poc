package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/kl-engineering/schedule/api/middleware/requestlogger"
	"github.com/kl-engineering/schedule/api/routes"
	"github.com/kl-engineering/schedule/pkg/backend/cmsclient"
	moretime "github.com/kl-engineering/schedule/pkg/util/time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
)

func main() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(moretime.RFC3339Milli)
	logger, _ := config.Build()
	defer logger.Sync()

	app := fiber.New()
	app.Use(
		requestid.New(),
		requestlogger.New(logger),
	)

	routes.PingRouter(app)

	client := cmsclient.New()
	routes.SchedulesRouter(app, client)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// allows for more graceful shutdown
	serverShutdown := make(chan struct{})
	go func() {
		_ = <-c
		logger.Info("Gracefully shutting down...")
		_ = app.Shutdown()
		serverShutdown <- struct{}{}
	}()

	if err := app.Listen(":3000"); err != nil {
		logger.Panic("", zap.Error(err))
	}

	<-serverShutdown
}

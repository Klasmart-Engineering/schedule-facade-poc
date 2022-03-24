package eventpublisher

import "go.uber.org/zap"

type Publisher interface {
	Publish(event interface{})
}

type publisher struct {
	logger *zap.Logger
}

func (p *publisher) Publish(event interface{}) {
	p.logger.Info("published", zap.Any("event", event))
}

func New(logger *zap.Logger) *publisher {
	return &publisher{
		logger: logger,
	}
}

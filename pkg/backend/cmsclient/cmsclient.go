package cmsclient

import (
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Client interface {
	GetSchedule(id string) (string, error)
}

type client struct {
	logger *zap.Logger
}

func (c *client) GetSchedule(id string) (string, error) {
	// TODO(JR): authentication?
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8088/v1/schedules/"+id, nil)
	if err != nil {
		c.logger.Error("error", zap.Error(err))
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		c.logger.Error("error", zap.Error(err))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("error", zap.Error(err))
	}

	return string(body), nil
}

func New(logger *zap.Logger) *client {
	return &client{
		logger: logger,
	}
}

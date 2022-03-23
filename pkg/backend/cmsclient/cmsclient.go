package cmsclient

type Client interface {
	GetSchedule(id string) (string, error)
}

type client struct {
}

func (c *client) GetSchedule(id string) (string, error) {
	return "{}", nil
}

func New() *client {
	return &client{}
}

package status

import (
	"context"
	"net/http"

	"github.com/IAmRasputin/artifacts-go/internal/client"
)

type GameStatusClient interface {
	GetGameServerStatus() (*http.Response, error)
}

type Client struct {
	ctx            context.Context
	internalClient *client.Client
}

func NewGameStatusClient(artifactsClient *client.Client) GameStatusClient {
	return &Client{
		ctx:            context.Background(),
		internalClient: artifactsClient,
	}
}

func (c *Client) GetGameServerStatus() (*http.Response, error) {
	resp, err := c.internalClient.GetStatusGet(c.ctx, client.IncludeAuth)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

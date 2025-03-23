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
	internalClient GameStatusGetter
}

type GameStatusGetter interface {
	GetStatusGet(context.Context, ...client.RequestEditorFn) (*http.Response, error)
}

func NewGameStatusClient(artifactsClient GameStatusGetter) GameStatusClient {
	return &Client{
		ctx:            context.Background(),
		internalClient: artifactsClient,
	}
}

func (c *Client) GetGameServerStatus() (*http.Response, error) {
	resp, err := c.internalClient.GetStatusGet(c.ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

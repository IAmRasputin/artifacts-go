package status

import (
	"context"

	"github.com/IAmRasputin/artifacts-go/internal/client"
)

type GameStatusClient interface {
	GetGameServerStatus() (*client.GetStatusGetResponse, error)
}

type Client struct {
	ctx            context.Context
	internalClient GameStatusGetter
}

type GameStatusGetter interface {
	GetStatusGetWithResponse(context.Context, ...client.RequestEditorFn) (*client.GetStatusGetResponse, error)
}

func NewGameStatusClient(artifactsClient GameStatusGetter) GameStatusClient {
	return &Client{
		ctx:            context.Background(),
		internalClient: artifactsClient,
	}
}

func (c *Client) GetGameServerStatus() (*client.GetStatusGetResponse, error) {
	resp, err := c.internalClient.GetStatusGetWithResponse(c.ctx)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

package status

import (
	"context"

	"github.com/IAmRasputin/artifacts-go/internal/client"
)

type GameStatus = client.StatusSchema

type GameStatusClient interface {
	GetGameServerStatus() (*GameStatus, error)
}

type GameStatusGetter interface {
	GetStatusGetWithResponse(context.Context, ...client.RequestEditorFn) (*client.GetStatusGetResponse, error)
}

type artifactsClient struct {
	ctx            context.Context
	internalClient GameStatusGetter
}

func NewGameStatusClient(afClient GameStatusGetter) GameStatusClient {
	return &artifactsClient{
		ctx:            context.Background(),
		internalClient: afClient,
	}
}

func NewDefaultGameStatusClient() (GameStatusClient, error) {
	artifactsClient, err := client.NewClientWithResponses(client.BaseURL)

	if err != nil {
		return nil, err
	}

	return NewGameStatusClient(artifactsClient), nil
}

func (c *artifactsClient) GetGameServerStatus() (*GameStatus, error) {
	resp, err := c.internalClient.GetStatusGetWithResponse(c.ctx)

	if err != nil {
		return nil, err
	}

	return &resp.JSON200.Data, nil
}

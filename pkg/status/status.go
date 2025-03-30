package status

import (
	"context"

	"github.com/IAmRasputin/artifacts-go/internal/client"
)

type GameStatus = client.StatusSchema

type GameStatusClient interface {
	GetGameServerStatus() (*GameStatus, error)
}

type artifactsClient struct {
	ctx            context.Context
	internalClient *client.ClientWithResponses
}

func NewGameStatusClient(afClient *client.ClientWithResponses) (GameStatusClient, error) {
	return &artifactsClient{
		ctx:            context.Background(),
		internalClient: afClient,
	}, nil
}

func NewDefaultGameStatusClient() (GameStatusClient, error) {
	afClient, err := client.NewClientWithResponses(client.BaseURL)

	if err != nil {
		return nil, err
	}

	return &artifactsClient{
		ctx:            context.Background(),
		internalClient: afClient,
	}, nil
}

func (c *artifactsClient) GetGameServerStatus() (*GameStatus, error) {
	resp, err := c.internalClient.GetStatusGetWithResponse(c.ctx, client.DefaultAuth)

	if err != nil {
		return nil, err
	}

	return &resp.JSON200.Data, nil
}

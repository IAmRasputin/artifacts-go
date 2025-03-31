package characters

import (
	"context"

	"github.com/IAmRasputin/artifacts-go/internal/client"
)

type Character = client.CharacterSchema
type Log = client.LogSchema

type CharacterClient interface {
	GetCharacters() ([]Character, error)
	GetCharacterLogs() ([]Log, error)
}

type artifactsClient struct {
	ctx            context.Context
	internalClient *client.ClientWithResponses
}

func NewCharacterClient(afClient *client.ClientWithResponses) CharacterClient {
	return &artifactsClient{
		ctx:            context.Background(),
		internalClient: afClient,
	}
}

func NewDefaultCharacterClient() (CharacterClient, error) {
	afClient, err := client.NewClientWithResponses(client.BaseURL)

	if err != nil {
		return nil, err
	}

	return NewCharacterClient(afClient), nil
}

func (a *artifactsClient) GetCharacters() ([]Character, error) {
	resp, err := a.internalClient.GetMyCharactersMyCharactersGetWithResponse(a.ctx, client.DefaultAuth)

	if err != nil {
		return nil, err
	}

	return resp.JSON200.Data, nil
}

func (a *artifactsClient) GetCharacterLogs() ([]Log, error) {
	return []Log{}, nil
}

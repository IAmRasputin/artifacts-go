package pkg

import (
	"os"

	"github.com/IAmRasputin/artifacts-go/internal/client"
	"github.com/IAmRasputin/artifacts-go/pkg/status"
)

type Artifacts struct {
	gameStatusClient status.GameStatusClient
}

func NewArtifacts() *Artifacts {
	gameClient, err := client.NewClient(client.BaseURL)
	if err != nil {
		// There's probably no way to recover from this, perish immediately
		os.Exit(1)
	}

	return &Artifacts{
		gameStatusClient: status.NewGameStatusClient(gameClient),
	}
}

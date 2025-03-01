package api

import (
	"net/http"

	"github.com/IAmRasputin/artifacts-go/internal/config"
)

type ArtifactsClient struct {
	http.Client
	tokenProvider config.TokenGetter
}

func NewArtifactsClient() *ArtifactsClient {
	// Should this return an error?  Maybe, but if it fails, we just want
	// to die instantly
	return &ArtifactsClient{
		Client:        http.Client{}, // Other options go here
		tokenProvider: config.NewDefaultTokenGetter(),
	}
}

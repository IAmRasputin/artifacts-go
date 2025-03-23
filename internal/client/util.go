package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IAmRasputin/artifacts-go/internal/config"
)

const BaseURL string = "https://api.artifactsmmo.com"

func DefaultAuth(ctx context.Context, req *http.Request) error {
	defaultToken := config.NewDefaultTokenGetter()

	return IncludeAuth(ctx, req, defaultToken)
}

func IncludeAuth(ctx context.Context, req *http.Request, tokenGetter config.TokenGetter) error {
	tok, err := tokenGetter.GetToken()

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tok))

	return nil
}

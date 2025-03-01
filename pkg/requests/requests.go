package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/IAmRasputin/artifacts-go/internal/config"
)

func RequestWithToken(method, path string, tokenProvider config.TokenGetter) (*http.Request, error) {
	fullUrl := url.URL{
		Scheme: "https",
		Host:   "api.artifactsmmo.com",
		Path:   path,
	}

	req, err := http.NewRequest(method, fullUrl.String(), strings.NewReader(""))

	if err != nil {
		return nil, err
	}

	token, err := tokenProvider.GetToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return req, nil
}

func RequestWithoutToken(method, path string) (*http.Request, error) {
	fullUrl := url.URL{
		Scheme: "https",
		Host:   "api.artifactsmmo.com",
		Path:   path,
	}

	req, err := http.NewRequest(method, fullUrl.String(), strings.NewReader(""))

	if err != nil {
		return nil, err
	}

	return req, nil
}

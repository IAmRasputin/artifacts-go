package requests

import (
	"reflect"
	"testing"
)

type TestTokenProvider struct{}

func (t *TestTokenProvider) GetToken() (string, error) {
	return "test_token", nil
}

func TestRequestWithToken(t *testing.T) {
	tokenProvider := TestTokenProvider{}
	req, err := RequestWithToken("GET", "my/path", &tokenProvider)

	if err != nil {
		t.Fatal(err)
	}

	if req.Host != "api.artifactsmmo.com" {
		t.Error("failed to set host")
	}

	if req.URL.Path != "/my/path" {
		t.Error("wrong path returned from request")
	}

	if !reflect.DeepEqual(req.Header["Authorization"], []string{"Bearer test_token"}) {
		t.Error("Failed to set auth header")
	}
}

func TestRequestWithoutToken(t *testing.T) {
	req, err := RequestWithoutToken("GET", "my/path")

	if err != nil {
		t.Fatal(err)
	}

	if req.Host != "api.artifactsmmo.com" {
		t.Error("failed to set host")
	}

	if req.URL.Path != "/my/path" {
		t.Error("path is not what it should be")
	}
}

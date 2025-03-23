package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

type TestTokenGetter struct {
	token string
	err   error
}

func (t *TestTokenGetter) GetToken() (string, error) {
	if t.err != nil {
		return "", t.err
	}

	return t.token, nil
}

func TestIncludeAuth(t *testing.T) {
	ctx := context.TODO()

	req := http.Request{}
	tokenGetter := &TestTokenGetter{
		token: "test_token",
		err:   nil,
	}

	err := IncludeAuth(ctx, &req, tokenGetter)

	if err != nil {
		t.Fatal("?")
	} else {
		if req.Header["Authorization"][0] != fmt.Sprintf("Bearer %s", tokenGetter.token) {
			t.Error("got wrong token from token getter")
		}
	}
}

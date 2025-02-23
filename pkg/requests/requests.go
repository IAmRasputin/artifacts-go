package requests

import (
	"net/http"
	"net/url"
)

func RequestWithToken(method, path string) (*http.Request, error) {
	u := url.URL{}
}

package dummyclient

import (
	"context"
	"fmt"
	"konntent-workspace-service/pkg/constants"
	"konntent-workspace-service/pkg/dummyclient/model"
	"konntent-workspace-service/pkg/httpclient"
)

type Client interface {
	HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error)

	PrepareHeaders(apiKey string) map[string]string
	PrepareURL(uri string) string

	Dummy(ctx context.Context, r model.DummyRequest) error
}

type Config struct {
	URL string
}

type client struct {
	BaseURL    string
	httpClient httpclient.HTTPClient
}

func NewClient(c Config, hc httpclient.HTTPClient) Client {
	return &client{
		BaseURL:    c.URL,
		httpClient: hc,
	}
}

func (c *client) HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error) {
	resp, err := c.httpClient.HandleRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if c.httpClient.IsSuccessStatusCode(resp) {
		return resp, nil
	}

	return nil, c.httpClient.HandleException(resp)
}

func (c *client) PrepareURL(uri string) string {
	return fmt.Sprintf("%s/%s", c.BaseURL, uri)
}

func (c *client) PrepareHeaders(apiKey string) map[string]string {
	return map[string]string{
		constants.HeaderKeyContentType:   "application/json",
		constants.HeaderKeyAuthorization: "Basic " + apiKey,
	}
}

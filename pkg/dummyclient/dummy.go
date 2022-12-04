package dummyclient

import (
	"context"
	"konntent-workspace-service/pkg/dummyclient/model"
	"konntent-workspace-service/pkg/httpclient"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	toMessageURL = "text/advanced"
)

func (c *client) Dummy(ctx context.Context, r model.DummyRequest) error {
	resp, err := c.HandleRequest(ctx, httpclient.Request{
		URL:    c.PrepareURL(toMessageURL),
		Method: fiber.MethodGet,
		Body:   r,
		//Headers: c.PrepareHeaders(middlewarepkg.GetHeaders(ctx)),
		Timeout: time.Second * 5,
	})

	log.Println(resp)
	return err
}

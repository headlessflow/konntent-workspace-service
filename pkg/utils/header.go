package utils

import "context"

type ContextHeader struct {
	UserID string `reqHeader:"X-User-ID"`
}

func GetHeaderMapByContext(c context.Context) *ContextHeader {
	return c.Value("headers").(*ContextHeader)
}

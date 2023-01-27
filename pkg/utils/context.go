package utils

import (
	"context"
	"strconv"
)

func GetUserIDByContext(c context.Context) int {
	var headers = GetHeaderMapByContext(c)
	if _, ok := headers["X-User-ID"]; !ok {
		return 0
	}

	var uid, err = strconv.Atoi(headers["X-User-ID"])
	if err != nil {
		return 0
	}

	return uid
}

func GetHeaderMapByContext(c context.Context) map[string]string {
	return c.Value("headers").(map[string]string)
}

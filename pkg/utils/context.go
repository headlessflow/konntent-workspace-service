package utils

import (
	"context"
	"strconv"
)

func GetUserIDByContext(c context.Context) int {
	var headers = GetHeaderMapByContext(c)

	var uid, err = strconv.Atoi(headers.UserID)
	if err != nil {
		return 0
	}

	return uid
}

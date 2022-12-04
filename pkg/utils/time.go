package utils

import (
	"konntent-workspace-service/pkg/constants"
	"time"
)

func PreCalculateTimeDiff(t1 time.Time, t2 string) int64 {
	if t2 == "" {
		return 0
	}

	t2Time, parseErr := time.Parse(constants.ParseLayout, t2)
	if parseErr != nil {
		return 0
	}

	return t2Time.Sub(t1).Milliseconds()
}

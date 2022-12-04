package utils

import "github.com/streadway/amqp"

const (
	XDeathKeyName      = "x-death"
	XDeathCountKeyName = "count"
)

func GetXDeathCount(m map[string]interface{}) int {
	death, ok := m[XDeathKeyName].([]interface{})
	if !ok {
		return 0
	}

	return int(death[0].(amqp.Table)[XDeathCountKeyName].(int64))
}

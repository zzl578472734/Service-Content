package utils

import (
	"Service-Content/constants"
	"time"
)

func TimeFormat(t time.Time) string {
	return t.Format(constants.DefaultLayout)
}

/**
 * 验证数字number是否在min和max里面
 */
func IntRange(number, min, max int) bool {
	if number < min || number > max {
		return false
	}
	return true
}

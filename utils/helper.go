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

/**
 * 处理分页
 */
func FilterPage(param int)  int{
	var page = 0
	if param < page{
		return page
	}
	return param
}

/**
 * 处理分页大小
 */
func FilterPageSize(param int)  int{
	var minPageSize = 0
	var maxPageSize = 100000
	var defaultPageSize  = 10
	if param < minPageSize || param > maxPageSize{
		param = defaultPageSize
	}
	return param
}
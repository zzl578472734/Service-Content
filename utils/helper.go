package utils

import (
	"Service-Content/constants"
	"time"
	"math/rand"
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

/**
 * 生成随机字符串
 */
func  GetRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randStr := make([]rune, length)
	for i := range randStr {
		randStr[i] = letters[rand.Intn(len(letters))]
	}
	return string(randStr)
}

/**
 * 验证数字number是小于min
 */
func IntMin(number, min int) bool {
	if number < min {
		return false
	}
	return true
}
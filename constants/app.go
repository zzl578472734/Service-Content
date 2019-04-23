package constants

import "time"

const (
	DefaultZero = 0
	DefaultEmptyString = ""

	DefaultApiSuccessCode = 0
	DefaultApiSuccessMsg = "success"

	DefaultLayout = "2016-01-02 15:04:05"
	DefaultErrorTemplate = "%v: %v error, detail: %v"

	DefaultCacheExpire = 10 * time.Minute
)
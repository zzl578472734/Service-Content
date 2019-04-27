package constants

import "time"

const (
	DefaultZero = 0
	DefaultEmptyString = ""

	DefaultApiSuccessCode = 0
	DefaultApiSuccessMsg = "success"

	DefaultLayout = "2006-01-02 15:04:05"
	DefaultErrorTemplate = "%v: %v error, detail: %v"
	DefaultRequestMaxTimestamps = 5 * time.Minute
	DefaultRequestMinTimestamps = -5 * time.Minute

	SearchIndex = "cym_content"

)
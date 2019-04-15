package utils

import (
	"time"
	"Service-Content/constants"
)

func FormatTime(t time.Time)  string{
	return t.Format(constants.DefaultLayout)
}
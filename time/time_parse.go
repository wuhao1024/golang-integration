package time

import (
	"time"
)

// ParseStringToTime parse string to time 基础方法
func ParseStringToTime(timeStr, layout string) time.Time {
	t, _ := time.ParseInLocation(layout, timeStr, LocalTime)
	return t
}

// ParseStringToDate return "2021-01-01 00:00:00 +0800 CST"(local time)
func ParseStringToDate(timeStr string) time.Time {
	return ParseStringToTime(timeStr, DefaultDateFormat)
}

// ParseStringToDateTime return "2021-01-01 01:01:01 +0800 CST"(local time)
func ParseStringToDateTime(timeStr string) time.Time {
	return ParseStringToTime(timeStr, DefaultDateTimeFormat)
}

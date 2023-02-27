package time

import (
	"time"
)

// FormatTime 基础方法
// 依据layout的格式返回
// DefaultDateTimeFormat  2023-02-27 16:11:06
// DefaultDateFormat  2023-02-27
func FormatTime(t time.Time, layout string) string {
	return t.In(LocalTime).Format(layout)
}

// FormatDate return "2006-01-02" format string
func FormatDate(t time.Time) string {
	return FormatTime(t, DefaultDateFormat)
}

// FormatDateTime return "2006-01-02 15:04:05" format string
func FormatDateTime(t time.Time) string {
	return FormatTime(t, DefaultDateTimeFormat)
}

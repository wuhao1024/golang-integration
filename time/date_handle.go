package time

import (
	"time"
)

// AddDate 时间增减 +为增 -为减
// 如：t 为 2020-03-30 12:00:00 +0800，增加1个月，为4月30号
// 如：t 为 2020-03-31 12:00:00 +0800，增加1个月，为4月30号
func AddDate(t time.Time, years int, months int, days int) time.Time {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()

	// firstDayOfMonthAfterAddDate: years年 & months月之后的 那个月份的 1号
	firstDayOfMonthAfterAddDate := time.Date(year+years, month+time.Month(months), 1, hour, min, sec, t.Nanosecond(), t.Location())

	// firstDayOfMonthAfterAddDate 月份的最后一天
	lastDay := LastDayOfMonth(firstDayOfMonthAfterAddDate)

	if day > lastDay {
		day = lastDay
	}

	return time.Date(year+years, month+time.Month(months), day+days, hour, min, sec, t.Nanosecond(), t.Location())
}

// LastDayOfMonth returns the last day of the month specified by local time of t.
func LastDayOfMonth(t time.Time) int {
	return EndOfMonth(t).Day()
}

// EndOfMonth end of month(local time)
func EndOfMonth(t time.Time) time.Time {
	return BeginningOfMonth(t).AddDate(0, 1, 0).Add(time.Microsecond)
}

// BeginningOfMonth beginning of month(local time)
func BeginningOfMonth(t time.Time) time.Time {
	year, month, _ := t.In(LocalTime).Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, LocalTime)
}

package time

import "time"

var DefaultDateTimeFormat = "2006-01-02 15:04:05"
var DefaultDateFormat = "2006-01-02"
var LocalTime, _ = time.LoadLocation("Asia/Shanghai")

// MysqlMaxTimeNanoseconds mysql datetime(6) 字段操作最大的纳秒数(超过此值999_999_499秒数会进位)
var MysqlMaxTimeNanoseconds = 999_999_000

var EndOf20991231 = time.Date(2099, 12, 31, 23, 59, 59, MysqlMaxTimeNanoseconds, LocalTime)

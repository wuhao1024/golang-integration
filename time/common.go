package time

import "time"

var DefaultDateTimeFormat = "2006-01-02 15:04:05"
var DefaultDateFormat = "2006-01-02"
var LocalTime, _ = time.LoadLocation("Asia/Shanghai")

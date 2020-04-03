package utils

import (
	"fmt"
	"time"
)

func YearMonthDay(byTime int64) string {
	t := time.Unix(byTime, 0)
	formatted := fmt.Sprintf("%02d.%02d.%d",
		t.Day(), t.Month(), t.Year())

	return formatted
}

func MonthDayWeekName(byTime int64) string {
	t := time.Unix(byTime, 0)
	formatted := fmt.Sprintf("%02d.%02d %02s",
		t.Month(), t.Day(), t.Weekday())

	return formatted
}

func ResolveByTimezone(tz int) int64 {
	return int64((tz - 3) * 3600)
}

func AsTime(byTime int64) string {
	t := time.Unix(byTime, 0)

	formatted := fmt.Sprintf("%02d-%02d",
		t.Hour(), t.Minute())

	return formatted
}

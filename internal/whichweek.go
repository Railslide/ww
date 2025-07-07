package internal

import (
	"time"
)

type WeekDateRange struct {
	Start time.Time
	End   time.Time
}

func CalculateWeek(week, year int) WeekDateRange {
	// Find the week of January 1st
	janFirst := time.Date(year, time.January, 0, 0, 0, 0, 0, time.UTC)
	monday := janFirst.AddDate(0, 0, 1 - int(janFirst.Weekday()))
	sunday := janFirst.AddDate(0, 0, 7 - int(janFirst.Weekday()))

	// Remove one week if January 1st is already in week one,
	// so that week one doesn't get counted twice.
	if _, w := janFirst.ISOWeek(); w == 1 {
		week = week - 1
	}

	// number of days between the week of January 1st and the target week
	days := 7 * week
	return WeekDateRange{
		Start: monday.AddDate(0, 0, days),
		End:   sunday.AddDate(0, 0, days),
	}
}

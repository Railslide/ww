package internal

import (
	"time"
)

type WeekDateRange struct {
	Start time.Time
	End   time.Time
}

func CalculateWeek(week, year int) WeekDateRange {
	janFirst := time.Date(year, time.January, 0, 0, 0, 0, 0, time.UTC) // TODO: change timezone?
	monday := janFirst.AddDate(0, 0, (1 - int(janFirst.Weekday())))
	sunday := janFirst.AddDate(0, 0, 7 - int(janFirst.Weekday()))

	// Amount of days between the first week of the year and the target one
	days := 7 * (week - 1)
	return WeekDateRange{
		Start: monday.AddDate(0, 0, days),
		End:   sunday.AddDate(0, 0, days),
	}
}

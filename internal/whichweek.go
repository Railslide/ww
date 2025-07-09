package internal

import (
	"errors"
	"time"
)

type WeekDateRange struct {
	Start time.Time
	End   time.Time
}

func CalculateWeek(week, year int) (WeekDateRange, error) {

	if week < 1 || week > 53 {
		return WeekDateRange{}, errors.New("Invalid week number")
	}

	// Find the week of January 1st
	janFirst := time.Date(year, time.January, 0, 0, 0, 0, 0, time.UTC)
	monday := janFirst.AddDate(0, 0, 1-int(janFirst.Weekday()))
	sunday := janFirst.AddDate(0, 0, 7-int(janFirst.Weekday()))

	// number of days between the beginning of the year and the target week
	days := 7 * week

	// Remove one week if January 1st is already in week one,
	// so that week one doesn't get counted twice.
	if _, w := janFirst.ISOWeek(); w == 1 {
		days = days - 7
	}

	wr := WeekDateRange{
		Start: monday.AddDate(0, 0, days),
		End:   sunday.AddDate(0, 0, days),
	}

	if _, w := wr.Start.ISOWeek(); w != week {
		// There is no week 53 for the given year
		if week == 53 && w == 1 {
			return WeekDateRange{}, errors.New("Invalid week number")
		} else {
			return WeekDateRange{}, errors.New("Something went wrong... :(")
		}
	}

	return wr, nil
}

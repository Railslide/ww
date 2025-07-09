package internal

import (
	"testing"
	"time"
)

func TestCalculateWeek(t *testing.T) {
	// Todo: add a test for leap year?
	tests := []struct {
		name          string
		year          int
		week          int
		expectedStart string
		expectedEnd   string
		wantError     bool
	}{
		{
			name:          "Basic case",
			year:          2025,
			week:          18,
			expectedStart: "2025-04-28",
			expectedEnd:   "2025-05-04",
		},
		{
			name:          "Jan 1 in week 53",
			year:          2027,
			week:          3,
			expectedStart: "2027-01-18",
			expectedEnd:   "2027-01-24",
		},
		{
			name:          "Jan 1 in week 52",
			year:          2023,
			week:          3,
			expectedStart: "2023-01-16",
			expectedEnd:   "2023-01-22",
		},
		{
			name: "Week number too high",
			year: 2025,
			week: 54,
			wantError: true,
		},
		{
			name: "Week number too low",
			year: 2025,
			week: -1,
			wantError: true,
		},
		{
			name: "Week 53 with a 52-weeks year",
			year: 2025,
			week: 53,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr, err := CalculateWeek(tt.week, tt.year)

			if err != nil {
				if tt.wantError {
					return
				} else {
					t.Errorf("test got an unexpected error: %v", err)
				}
			} else {
				if tt.wantError {
					t.Error("an error was expected but no error was raised")
					return
				}
			}

			start, err := time.Parse("2006-01-02", tt.expectedStart)
			if err != nil {
				t.Error("error while parsing start date...")

			}
			end, err := time.Parse("2006-01-02", tt.expectedEnd)
			if err != nil {
				t.Error("error while parsing end date...")
			}

			expectedWr := WeekDateRange{Start: start, End: end}
			if wr != expectedWr {
				t.Errorf("got %v, expected: %v", wr, expectedWr)
			}
		})
	}
}

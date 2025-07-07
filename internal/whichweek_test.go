package internal

import (
	"testing"
	"time"
)

func TestCalculateWeek(t *testing.T) {
	// Todo: add a test for leap year?
	tests := []struct {
		name string
		year int
		week int
		expectedStart string
		expectedEnd string
	}{
		{
			name: "Basic case",
			year: 2025,
			week: 18,
			expectedStart: "2025-04-28",
			expectedEnd: "2025-05-04",
		},
		{
			name: "Jan 1 in week 53",
			year: 2027,
			week: 3,
			expectedStart: "2027-01-18",
			expectedEnd: "2027-01-24",
		},
		{
			name: "Jan 1 in week 52",
			year: 2023,
			week: 3,
			expectedStart: "2023-01-16",
			expectedEnd: "2023-01-22",
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := CalculateWeek(tt.week, tt.year)

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

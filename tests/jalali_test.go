package tests

import (
	"mohsenbakhit/jalali"
	"testing"
	"time"
)

func TestNewJalali(t *testing.T) {
	jalali := jalali.NewJalali(1403, 10, 12, 4)

	if jalali.Year() != 1403 {
		t.Errorf("Expected year to be 2025, got %d", jalali.Year())
	}

	if jalali.Month() != 10 {
		t.Errorf("Expected month to be 1, got %d", jalali.Month())
	}

	if jalali.Day() != 12 {
		t.Errorf("Expected day to be 1, got %d", jalali.Day())
	}

	if jalali.Weekday() != "Charshanbe" {
		t.Errorf("Expected weekday to be Charshanbe, got %s", jalali.Weekday())
	}

	if jalali.MonthString() != "Dey" {
		t.Errorf("Expected month string to be Dey, got %s", jalali.MonthString())
	}

	t.Logf("%d-%d-%d is a %s of a %s mah", jalali.Year(), jalali.Month(), jalali.Day(), jalali.Weekday(), jalali.MonthString())
}

func TestToJalali(t *testing.T) {
	// Test cases with known Gregorian to Jalali conversions
	testCases := []struct {
		gregorianYear   int
		gregorianMonth  int
		gregorianDay    int
		expectedYear    int
		expectedMonth   int
		expectedDay     int
		expectedWeekday string
	}{
		{2024, 1, 1, 1402, 10, 11, "Doshanbe"},
		{2024, 3, 20, 1403, 1, 1, "Charshanbe"},
		{2024, 3, 19, 1402, 12, 29, "Seshanbe"},
		{2025, 1, 1, 1403, 10, 11, "Charshanbe"},
	}

	for _, tc := range testCases {
		gregorianDate := time.Date(tc.gregorianYear, time.Month(tc.gregorianMonth), tc.gregorianDay, 0, 0, 0, 0, time.UTC)
		jalali := jalali.ToJalali(gregorianDate)

		if jalali.Year() != tc.expectedYear {
			t.Errorf("Expected year to be %d, got %d for Gregorian date %v: Jalali date %v", tc.expectedYear, jalali.Year(), gregorianDate, jalali)
		}

		if jalali.Month() != tc.expectedMonth {
			t.Errorf("Expected month to be %d, got %d for Gregorian date %v: Jalali date %v", tc.expectedMonth, jalali.Month(), gregorianDate, jalali)
		}

		if jalali.Day() != tc.expectedDay {
			t.Errorf("Expected day to be %d, got %d for Gregorian date %v: Jalali date %v", tc.expectedDay, jalali.Day(), gregorianDate, jalali)
		}

		if jalali.Weekday() != tc.expectedWeekday {
			t.Errorf("Expected weekday to be %s, got %s for Gregorian date %v: Jalali date %v", tc.expectedWeekday, jalali.Weekday(), gregorianDate, jalali)
		}
	}

}

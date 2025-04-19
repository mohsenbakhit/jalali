package jalali

import (
	"time"
)

func NewJalali(year int, month int, day int, weekday int) *Jalali {
	return &Jalali{
		day:     Day(day),
		weekday: Weekday((weekday + 1) % 7),
		month:   Month(month),
		year:    Year(year),
	}
}

// ToJalali converts a time.Time object into a Jalali date
func ToJalali(t time.Time) *Jalali {
	gYear := t.Year()
	gMonth := int(t.Month())
	gDay := t.Day()

	// Determine the Gregorian year in which this Jalali year begins
	pivotYear := gYear
	// Farvardin 1 falls on March 21, or March 20 in a leap
	jNewYearDay := 21
	isPivotLeap := (pivotYear%4 == 0 && pivotYear%100 != 0) || (pivotYear%400 == 0)
	if isPivotLeap {
		jNewYearDay = 20
	}

	// If we're before that March date, the Jalali year started in the previous Gregorian year
	if gMonth < 3 || (gMonth == 3 && gDay < jNewYearDay) {
		pivotYear--
		isPivotLeap = (pivotYear%4 == 0 && pivotYear%100 != 0) || (pivotYear%400 == 0)
		jNewYearDay = 21
		if isPivotLeap {
			jNewYearDay = 20
		}
	}

	// Now set the Jalali year
	jYear := pivotYear - 621

	isJalaliLeap := (jYear%4 == 0 && jYear%100 != 0) || (jYear%400 == 0)

	// Build pivot-year cumulative days for calculating Nowruz
	pivotGregorianDays := [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	if isPivotLeap {
		for i := 2; i < len(pivotGregorianDays); i++ {
			pivotGregorianDays[i]++
		}
	}
	newYearDayOfYear := pivotGregorianDays[2] + jNewYearDay

	// Build actual cumulative days for the date's Gregorian year
	isActualLeap := (gYear%4 == 0 && gYear%100 != 0) || (gYear%400 == 0)
	actualGregorianDays := [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	if isActualLeap {
		for i := 2; i < len(actualGregorianDays); i++ {
			actualGregorianDays[i]++
		}
	}
	dayOfYear := actualGregorianDays[gMonth-1] + gDay

	daysSinceNewYear := dayOfYear - newYearDayOfYear
	if daysSinceNewYear < 0 {
		daysSinceNewYear += 365 + btoi(isJalaliLeap)
	}

	jalaliMonths := [...]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	if isJalaliLeap {
		jalaliMonths[11] = 30
	}

	jMonth := 1
	for daysSinceNewYear >= jalaliMonths[jMonth-1] {
		daysSinceNewYear -= jalaliMonths[jMonth-1]
		jMonth++
	}
	jDay := daysSinceNewYear + 1

	return NewJalali(int(jYear), jMonth, jDay, int(t.Weekday()+1))

}

// ToGregorian converts a Jalali date into a time.Time object
func ToGregorian(jYear int, jMonth int, jDay int, jWeekday int) *time.Time {
	isJalaliLeap := (jYear%4 == 0 && jYear%100 != 0) || (jYear%400 == 0)

	gYear := jYear + 621
	isGregorianLeap := (gYear%4 == 0 && gYear%100 != 0) || (gYear%400 == 0)

	jNewYearDay := 21
	if isGregorianLeap {
		jNewYearDay = 20
	}
	// Compute the day‑of‑year for Farvardin 1 in the pivot Gregorian year
	pivotGregorianDays := [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	if isGregorianLeap {
		for i := 2; i < len(pivotGregorianDays); i++ {
			pivotGregorianDays[i]++
		}
	}
	newYearDayOfYear := pivotGregorianDays[2] + jNewYearDay

	jalaliMonths := [...]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	if isJalaliLeap {
		jalaliMonths[11] = 30
	}

	// Calculate days passed in Jalali calendar
	daysPassed := jDay
	for i := 0; i < jMonth-1; i++ {
		daysPassed += jalaliMonths[i]
	}

	// Add days from start of Gregorian year to Jalali new year
	yearLen := 365 + btoi(isGregorianLeap)
	gDayOfYear := daysPassed + newYearDayOfYear - 1

	// If we overflow past the end of the Gregorian year
	if gDayOfYear >= yearLen {
		gDayOfYear -= yearLen
		gYear++
		isGregorianLeap = (gYear%4 == 0 && gYear%100 != 0) || (gYear%400 == 0)
		// if it lands exactly on the first day, ensure day 1
		if gDayOfYear == 0 {
			gDayOfYear = 1
		}
	}

	// Find Gregorian month and day
	gregorianMonths := [...]int{31, 28 + btoi(isGregorianLeap), 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	gMonth := 1
	for gDayOfYear > gregorianMonths[gMonth-1] {
		gDayOfYear -= gregorianMonths[gMonth-1]
		gMonth++
	}
	gDay := gDayOfYear

	result := time.Date(gYear, time.Month(gMonth), gDay, 0, 0, 0, 0, time.UTC)
	return &result
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

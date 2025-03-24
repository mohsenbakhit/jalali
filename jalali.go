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
	isGregorianLeap := (t.Year()%4 == 0 && t.Year()%100 != 0) || (t.Year()%400 == 0)

	gYear := t.Year()
	gMonth := int(t.Month())
	gDay := t.Day()

	var jMonth, jDay int
	jYear := gYear - 621

	jNewYearDay := 21
	if isGregorianLeap {
		jNewYearDay = 20
	}

	if gMonth < 3 || (gMonth == 3 && gDay < jNewYearDay) {
		jYear--
	}

	isJalaliLeap := (jYear%4 == 0 && jYear%100 != 0) || (jYear%400 == 0)

	gregorianDays := [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	if isGregorianLeap {
		for i := 2; i < len(gregorianDays); i++ {
			gregorianDays[i]++
		}
	}

	daysSinceNewYear := gregorianDays[gMonth-1] + gDay - jNewYearDay
	if daysSinceNewYear < 0 {
		daysSinceNewYear += 365 + btoi(isGregorianLeap)
	}

	jalaliMonths := [...]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	if isJalaliLeap {
		jalaliMonths[11] = 30
	}

	jMonth = 1
	for daysSinceNewYear >= jalaliMonths[jMonth-1] {
		daysSinceNewYear -= jalaliMonths[jMonth-1]
		jMonth++
	}
	jDay = daysSinceNewYear

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
	gDayOfYear := daysPassed + jNewYearDay - 1

	// If we've passed into next Gregorian year
	if gDayOfYear > 365+btoi(isGregorianLeap) {
		gDayOfYear -= 365 + btoi(isGregorianLeap)
		gYear++
		isGregorianLeap = (gYear%4 == 0 && gYear%100 != 0) || (gYear%400 == 0)
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

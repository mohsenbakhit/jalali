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

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

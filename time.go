package farsical

// An integer that specifies the day of the month: between 1 and 31
type Day int

// An integer that specifies the day of the week: shanbe = 1, jomee = 7
type Weekday int

var weekdays = [...]string{
	"Shanbe",
	"Yekshanbe",
	"Doshanbe",
	"Seshanbe",
	"Charshanbe",
	"Panjshanbe",
	"Jomee",
}

// An integer that specifies the month of the year: Farvardin = 1, Esfand = 12
type Month int

// A string specifying the month of the year
type MonthString string

var months = [...]string{
	"Farvarding",
	"Ordibehesht",
	"Khordad",
	"Tir",
	"Mordad",
	"Shahrivar",
	"Mehr",
	"Aban",
	"Azar",
	"Dey",
	"Bahman",
	"Esfand",
}

// An integer that specifies the year
type Year int

// An object that specifies a time
type Time struct {
	day     Day
	weekday Weekday
	month   Month
	year    Year
}

func (t *Time) Day() int {
	return int(t.day)
}

func (t *Time) Weekday() string {
	return weekdays[t.weekday-1]
}

func (t *Time) MonthString() string {
	return months[t.month-1]
}

func (t *Time) Month() int {
	return int(t.month)
}

func (t *Time) Year() int {
	return int(t.year)
}

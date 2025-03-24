package jalali

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
	"Farvardin",
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

// An object that specifies a Jalali date
type Jalali struct {
	day     Day
	weekday Weekday
	month   Month
	year    Year
}

// Returns the Day of the Jalali date
func (t *Jalali) Day() int {
	return int(t.day)
}

// Returns the Weekday of the Jalali date
func (t *Jalali) Weekday() string {
	return weekdays[t.weekday-1]
}

// Returns the name of the Month of the Jalali date
func (t *Jalali) MonthString() string {
	return months[t.month-1]
}

// Returns the number of the month of the Jalali date
func (t *Jalali) Month() int {
	return int(t.month)
}

// Returns the year of the month of the Farsi Jalali
func (t *Jalali) Year() int {
	return int(t.year)
}

package banktime

// BankTime assists in calculating processing days that meet the US Federal Reserve Banks processing days.
//
// For holidays falling on Saturday, Federal Reserve Banks and Branches will be open the preceding Friday.
// For holidays falling on Sunday, all Federal Reserve Banks and Branches will be closed the following Monday.
// ACH and FedWire payments are not processed on weekends or the following US holidays.
// https://www.frbservices.org/holidayschedules/
// All logic is based on ET(Eastern) time as defined by the Federal Reserve
// https://www.frbservices.org/operations/fedwire/fedwire_hours.html

import (
	"time"

	"github.com/rickar/cal"
)

// BankTime takes a time.Time with accessors for US bank Holidays and banking days.
type BankTime struct {
	time time.Time
	cal  *cal.Calendar
}

// NewBankTime creates a new BankDate with an instantiated calendar
func NewBankTime(t time.Time) *BankTime {
	c := cal.NewCalendar()
	cal.AddUsHolidays(c)
	c.Observed = cal.ObservedMonday
	est, _ = time.LoadLocation("America/New_York")
	t = t.In(est)
	bt := &BankTime{time: t, cal: c}

	return bt
}

// IsBankingDay returns true if the day is a banking day, false otherwise
func (bt *BankTime) IsBankingDay() bool {
	// if date is not a weekend and not a holiday it is banking day.
	if bt.IsWeekend() {
		return false
	}
	// and not a holiday
	if bt.cal.IsHoliday(bt.time) {
		return false
	}
	// and not a monday after a holiday
	if bt.time.Weekday() == time.Monday {
		sun := bt.time.AddDate(0, 0, -1)
		return !bt.cal.IsHoliday(sun)
	}

	return true
}

// AddBankingDay takes an integer for the number of valid banking days to add and returns a Time
func (bt *BankTime) AddBankingDay(d int) time.Time {
	bt.time = bt.time.AddDate(0, 0, d)
	if !bt.IsBankingDay() {
		return bt.AddBankingDay(1)
	}
	return bt.time
}

// IsWeekend reports whether the given date falls on a weekend.
func (bt *BankTime) IsWeekend() bool {
	day := bt.time.Weekday()
	return day == time.Saturday || day == time.Sunday
}

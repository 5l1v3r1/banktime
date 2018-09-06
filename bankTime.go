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
	"sync"
	"time"

	"github.com/rickar/cal"
)

var (
	// DefaultLocation is the *time.Location used when creating a BankTime instance.
	// By default America/New_York is populated here.
	DefaultLocation *time.Location
	setup           sync.Once
)

// BankTime takes a time.Time with accessors for US bank Holidays and banking days.
type BankTime struct {
	time time.Time
	cal  *cal.Calendar
}

// New returns a BankTime for an optional *time.Location
// If no *time.Location is provided then America/New_York is used.
func New(t time.Time, loc *time.Location) *BankTime {
	c := cal.NewCalendar()
	cal.AddUsHolidays(c)
	c.Observed = cal.ObservedMonday

	if loc == nil {
		setup.Do(func() {
			loc, _ = time.LoadLocation("America/New_York")
			DefaultLocation = loc
		})
		t = t.In(DefaultLocation)
	} else {
		t = t.In(loc)
	}

	return &BankTime{
		time: t,
		cal:  c,
	}
}

// NewBankTime creates a new BankTime with an instantiated calendar
// in America/New_York.
func NewBankTime(t time.Time) *BankTime {
	return New(t, nil)
}

// IsBankingDay checks the rules around holidays (i.e. weekends) to
// determine if the given day is a banking day.
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
